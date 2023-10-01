// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/site"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/verror"
)

func Site(rc *fasthttp.RequestCtx) {
	path := util.StringSplitAndTrim(string(rc.Request.URI().Path()), "/")
	action := "site"
	if len(path) > 0 {
		action += "." + strings.Join(path, ".")
	}
	ActSite(action, rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		redir, page, bc, err := site.Handle(path, as, ps)
		if err != nil {
			return "", err
		}
		if _, ok := page.(*verror.NotFound); ok {
			rc.Response.SetStatusCode(404)
		}
		if redir != "" {
			return redir, nil
		}
		return Render(rc, as, page, ps, bc...)
	})
}
