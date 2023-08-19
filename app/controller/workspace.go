package controller

import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/views"
	"github.com/valyala/fasthttp"
)

func Workspace(rc *fasthttp.RequestCtx) {
	Act("workspace", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		x, err := as.Services.NPN.Collection.List(&ps.Profile.ID)
		if err != nil {
			return "", err
		}

		ps.Data = x
		return Render(rc, as, &views.Workspace{}, ps)
	})
}

func Socket(rc *fasthttp.RequestCtx) {
	Act("socket", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return "", nil
	})
}
