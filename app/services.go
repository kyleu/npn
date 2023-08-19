package app

import (
	"context"

	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/lib/scripting"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/imprt"
	"github.com/kyleu/npn/app/request/search"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/socket"
	"github.com/kyleu/npn/app/user"
	"github.com/kyleu/npn/app/util"
)

type Services struct {
	Script *scripting.Service
	Socket *websocket.Service
	User   *user.Service
	NPN    *socket.Service
	// add your dependencies here
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	scr := scripting.NewService(st.Files, "scripts")
	ws := websocket.NewService(nil, nil, nil)

	multiuser := util.GetEnvBool("npn_multiuser", false)
	us := user.NewService()
	ss := session.NewService(multiuser, st.Files, logger)
	co := collection.NewService(multiuser, st.Files, logger)
	rq := request.NewService(multiuser, st.Files, logger)
	cl := call.NewService(ss, logger)
	sr := search.NewService(co, rq, logger)
	im := imprt.NewService(st.Files, logger)
	npn := socket.NewService(us, ss, co, rq, cl, sr, im, ws)

	return &Services{Script: scr, Socket: ws, User: us, NPN: npn}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
