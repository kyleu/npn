// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/user"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/verror"
)

func Options(rc *fasthttp.RequestCtx) {
	cutil.WriteCORS(rc)
	rc.SetStatusCode(fasthttp.StatusOK)
}

func NotFound(rc *fasthttp.RequestCtx) {
	Act("notfound", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		cutil.WriteCORS(rc)
		rc.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
		rc.SetStatusCode(fasthttp.StatusNotFound)
		path := string(rc.Request.URI().Path())
		ps.Logger.Warnf("%s %s returned [%d]", string(rc.Method()), path, fasthttp.StatusNotFound)
		if ps.Title == "" {
			ps.Title = "Page not found"
		}
		ps.Data = ps.Title
		bc := util.StringSplitAndTrim(string(rc.URI().Path()), "/")
		bc = append(bc, "Not Found")
		return Render(rc, as, &verror.NotFound{Path: path}, ps, bc...)
	})
}

func Unauthorized(rc *fasthttp.RequestCtx, reason string, accounts user.Accounts) func(as *app.State, ps *cutil.PageState) (string, error) {
	return func(as *app.State, ps *cutil.PageState) (string, error) {
		cutil.WriteCORS(rc)
		rc.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
		rc.SetStatusCode(fasthttp.StatusUnauthorized)
		path := string(rc.Request.URI().Path())
		ps.Logger.Warnf("%s %s returned [%d]", string(rc.Method()), path, fasthttp.StatusNotFound)
		bc := util.StringSplitAndTrim(string(rc.URI().Path()), "/")
		bc = append(bc, "Unauthorized")
		if ps.Title == "" {
			ps.Title = "Unauthorized"
		}
		ps.Data = ps.Title
		return Render(rc, as, &verror.Unauthorized{Path: path, Message: reason, Accounts: accounts}, ps, bc...)
	}
}
