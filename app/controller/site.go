// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"net/http"
	"strings"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/site"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/verror"
)

func Site(w http.ResponseWriter, r *http.Request) {
	path := util.StringSplitAndTrim(r.URL.Path, "/")
	action := "site"
	if len(path) > 0 {
		action += "." + strings.Join(path, ".")
	}
	ActSite(action, w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		redir, page, bc, err := site.Handle(path, as, ps)
		if err != nil {
			return "", err
		}
		if _, ok := page.(*verror.NotFound); ok {
			w.WriteHeader(http.StatusNotFound)
		}
		if redir != "" {
			return redir, nil
		}
		return Render(r, as, page, ps, bc...)
	})
}
