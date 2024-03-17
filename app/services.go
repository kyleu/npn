package app

import (
	"context"

	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/lib/scripting"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/imprt"
	"github.com/kyleu/npn/app/request/search"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/socket"
	"github.com/kyleu/npn/app/util"
)

type Services struct {
	CoreServices

	Script *scripting.Service
	NPN    *socket.Service
	// add your dependencies here
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	core := initCoreServices(ctx, st, logger)

	multiuser := util.GetEnvBool("npn_multiuser", false)
	ss := session.NewService(multiuser, st.Files, logger)
	co := collection.NewService(multiuser, st.Files, logger)
	rq := request.NewService(multiuser, st.Files, logger)
	cl := call.NewService(ss, logger)
	sr := search.NewService(co, rq, logger)
	im := imprt.NewService(st.Files, logger)
	npn := socket.NewService(core.User, ss, co, rq, cl, sr, im)
	core.Socket = npn.Socket

	return &Services{CoreServices: core, NPN: npn}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
