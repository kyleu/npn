package app

import (
	"context"

	"github.com/kyleu/npn/app/lib/scripting"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request/transform"
	"github.com/kyleu/npn/app/util"
)

type Services struct {
	Script *scripting.Service
	Socket *websocket.Service
	// add your dependencies here
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	// TODO Remove
	_ = transform.RequestTransformers{}
	return &Services{
		Script: scripting.NewService(st.Files, "scripts"),
		Socket: websocket.NewService(nil, nil, nil),
	}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
