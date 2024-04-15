// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/telemetry"
	"github.com/kyleu/npn/app/lib/user"
	"github.com/kyleu/npn/app/site"
	"github.com/kyleu/npn/app/util"
)

type ActFn func(as *app.State, ps *cutil.PageState) (string, error)

func Act(key string, w http.ResponseWriter, r *http.Request, f ActFn) {
	as := _currentAppState
	wc := cutil.NewWriteCounter(w)
	ps := cutil.LoadPageState(as, wc, r, key, _currentAppRootLogger)
	if err := initAppRequest(as, ps); err != nil {
		ps.Logger.Warnf("%+v", err)
	}
	if !ps.Admin {
		if allowed, reason := user.Check(r.URL.Path, ps.Accounts); !allowed {
			f = Unauthorized(w, r, reason, ps.Accounts)
		}
	}
	actComplete(key, as, ps, wc, r, f)
}

func ActSite(key string, w http.ResponseWriter, r *http.Request, f func(as *app.State, ps *cutil.PageState) (string, error)) {
	as := _currentSiteState
	wc := cutil.NewWriteCounter(w)
	ps := cutil.LoadPageState(as, wc, r, key, _currentAppRootLogger)
	ps.Menu = site.Menu(ps.Context, as, ps.Profile, ps.Accounts, ps.Logger)
	if allowed, reason := user.Check(r.URL.Path, ps.Accounts); !allowed {
		f = Unauthorized(w, r, reason, ps.Accounts)
	}
	if err := initSiteRequest(as, ps); err != nil {
		ps.Logger.Warnf("%+v", err)
	}
	actComplete(key, as, ps, ps.W, r, f)
}

func actComplete(key string, as *app.State, ps *cutil.PageState, w *cutil.WriteCounter, r *http.Request, f ActFn) {
	err := ps.Clean(r, as)
	if err != nil {
		ps.Logger.Warnf("error while cleaning request, somehow: %+v", err)
	}
	status := http.StatusOK
	cutil.WriteCORS(w)
	var redir string
	logger := ps.Logger
	ctx := ps.Context
	if !telemetry.SkipControllerMetrics {
		var span *telemetry.Span
		ctx, span, logger = telemetry.StartSpan(ps.Context, "controller."+key, ps.Logger)
		defer span.Complete()
	}
	logger = logger.With("path", r.URL.Path, "method", ps.Method, "status", status)
	ps.Context = ctx

	if ps.ForceRedirect == "" || ps.ForceRedirect == r.URL.Path {
		redir, err = safeRun(f, as, ps)
		if err != nil {
			redir, err = handleError(key, as, ps, r, err)
			if err != nil {
				ps.Logger.Warnf("unable to handle error: %+v", err)
			}
		}
	} else {
		redir = ps.ForceRedirect
	}
	if redir != "" {
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
	}
	elapsedMillis := float64((util.TimeCurrentNanos()-ps.Started.UnixNano())/int64(time.Microsecond)) / float64(1000)
	ps.ResponseBytes = w.Count()
	defer ps.Close()
	w.Header().Set("Server-Timing", fmt.Sprintf("server:dur=%.3f", elapsedMillis))
	logger = logger.With("elapsed", elapsedMillis)
	logger.Debugf("processed request in [%.3fms] (render: %.3fms, response: %s)", elapsedMillis, ps.RenderElapsed, util.ByteSizeSI(ps.ResponseBytes))
}

func safeRun(f func(as *app.State, ps *cutil.PageState) (string, error), as *app.State, ps *cutil.PageState) (s string, e error) {
	defer func() {
		if rec := recover(); rec != nil {
			if recoverErr, ok := rec.(error); ok {
				e = errors.Wrap(recoverErr, "panic")
			} else {
				e = errors.Errorf("controller encountered panic recovery of type [%T]: %s", rec, fmt.Sprint(rec))
			}
		}
	}()
	s, e = f(as, ps)
	return
}
