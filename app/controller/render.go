package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views"
	"github.com/kyleu/npn/views/layout"
	"github.com/kyleu/npn/views/verror"
)

func Render(r *http.Request, as *app.State, page layout.Page, ps *cutil.PageState, breadcrumbs ...string) (string, error) {
	defer func() {
		x := recover()
		if x != nil {
			ps.LogError("error processing template: %+v", x)
			ps.W.WriteHeader(http.StatusInternalServerError)
			var ed *util.ErrorDetail
			switch t := x.(type) {
			case error:
				ed = util.GetErrorDetail(t, ps.Admin)
			default:
				ed = &util.ErrorDetail{Type: fmt.Sprintf("%T", x), Message: fmt.Sprint(t)}
			}
			if t := fmt.Sprintf("%T", page); t != "*verror.Error" {
				ps.Data = ed
				_, err := Render(r, as, &verror.Error{Err: ed}, ps, breadcrumbs...)
				if err != nil {
					verror.WriteDetail(ps.W, ed, as, ps)
				}
			} else {
				verror.WriteDetail(ps.W, ed, as, ps)
			}
		}
	}()

	var fn string
	if r.URL.Query().Get("download") == "true" {
		fn = ps.Action
	}

	maybeRender := func(ct string) (bool, error) {
		switch {
		case cutil.IsContentTypeCSV(ct):
			_, err := cutil.RespondCSV(ps.W, fn, ps.Data)
			return true, err
		case cutil.IsContentTypeJSON(ct):
			_, err := cutil.RespondJSON(ps.W, fn, ps.Data)
			return true, err
		case cutil.IsContentTypeTOML(ct):
			_, err := cutil.RespondTOML(ps.W, ps.Action, ps.Data)
			return true, err
		case cutil.IsContentTypeXML(ct):
			_, err := cutil.RespondXML(ps.W, fn, ps.Data)
			return true, err
		case cutil.IsContentTypeYAML(ct):
			_, err := cutil.RespondYAML(ps.W, fn, ps.Data)
			return true, err
		case cutil.IsContentTypeDebug(ct):
			_, err := cutil.RespondDebug(ps.W, r, as, fn, ps)
			return true, err
		}
		return false, nil
	}

	ps.Breadcrumbs = append(ps.Breadcrumbs, breadcrumbs...)
	ct := cutil.GetContentType(r)
	if ps.Data != nil {
		if handled, err := maybeRender(ct); handled || err != nil {
			return "", err
		}
	}
	startNanos := util.TimeCurrentNanos()
	if ps.DefaultFormat == "" {
		ps.W.Header().Set(cutil.HeaderContentType, "text/html; charset=UTF-8")
		views.WriteRender(ps.W, page, as, ps)
		ps.RenderElapsed = float64((util.TimeCurrentNanos()-startNanos)/int64(time.Microsecond)) / float64(1000)
		return "", nil
	}
	if handled, err := maybeRender(ps.DefaultFormat); handled || err != nil {
		return "", err
	}
	return "", errors.Errorf("unable to process format [%s]", ps.DefaultFormat)
}
