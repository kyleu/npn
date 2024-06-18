package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller"
	"github.com/kyleu/npn/app/controller/clib"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
)

func makeRoute(x *mux.Router, method string, path string, f http.HandlerFunc) {
	cutil.AddRoute(method, path)
	x.HandleFunc(path, f).Methods(method)
}

//nolint:revive
func AppRoutes(as *app.State, logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, "/", controller.Home)
	makeRoute(r, http.MethodGet, "/healthcheck", clib.Healthcheck)
	makeRoute(r, http.MethodGet, "/about", clib.About)

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.Profile)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)
	makeRoute(r, http.MethodGet, "/auth/{key}", clib.AuthDetail)
	makeRoute(r, http.MethodGet, "/auth/callback/{key}", clib.AuthCallback)
	makeRoute(r, http.MethodGet, "/auth/logout/{key}", clib.AuthLogout)
	makeRoute(r, http.MethodGet, cutil.DefaultSearchPath, clib.Search)

	themeRoutes(r)

	// $PF_SECTION_START(routes)$
	makeRoute(r, http.MethodGet, "/x", controller.Workspace)
	makeRoute(r, http.MethodGet, "/a", controller.Workspace)
	makeRoute(r, http.MethodGet, "/cfg", controller.Workspace)
	makeRoute(r, http.MethodGet, "/help", controller.Workspace)
	makeRoute(r, http.MethodGet, "/u", controller.Workspace)
	makeRoute(r, http.MethodGet, "/s", controller.Workspace)
	makeRoute(r, http.MethodGet, "/s/{extra:.*}", controller.Workspace)
	makeRoute(r, http.MethodGet, "/c", controller.Workspace)
	makeRoute(r, http.MethodGet, "/c/{extra:.*}", controller.Workspace)
	makeRoute(r, http.MethodGet, "/r", controller.Workspace)
	makeRoute(r, http.MethodGet, "/svg/gantt", controller.Gantt)
	makeRoute(r, http.MethodGet, "/ws", controller.Socket)
	// $PF_SECTION_END(routes)$

	scriptingRoutes(r)
	adminRoutes(r)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/robots.txt", clib.RobotsTxt)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodOptions, "/", controller.Options)

	return cutil.WireRouter(r, controller.NotFoundAction, logger)
}
