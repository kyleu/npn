package controller

import (
	"net/http"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views"
)

var homeContent = util.ValueMap{
	util.AppKey: util.AppName,
	"urls": map[string]string{
		"home": "/",
	},
}

func Home(w http.ResponseWriter, r *http.Request) {
	Act("home", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = homeContent
		return Render(r, as, &views.Workspace{}, ps)
	})
}
