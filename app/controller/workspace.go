package controller

import (
	"net/http"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views"
)

func Workspace(w http.ResponseWriter, r *http.Request) {
	Act("workspace", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		x, err := as.Services.NPN.Collection.List(&ps.Profile.ID)
		if err != nil {
			return "", err
		}
		ps.Data = x
		return Render(r, as, &views.Workspace{}, ps)
	})
}

func Socket(w http.ResponseWriter, r *http.Request) {
	Act("socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ch := util.RandomString(16)
		err := as.Services.Socket.Upgrade(ps.Context, w, r, ch, ps.User, ps.Profile, ps.Accounts, ps.Logger)
		if err != nil {
			ps.Logger.Warnf("unable to upgrade connection to WebSocket: %s", err.Error())
			return "", err
		}
		return "", nil
	})
}
