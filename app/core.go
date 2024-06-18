package app

import (
	"context"

	"github.com/kyleu/npn/app/lib/scripting"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/user"
	"github.com/kyleu/npn/app/util"
)

type CoreServices struct {
	User   *user.Service
	Script *scripting.Service
	Socket *websocket.Service
}

//nolint:revive
func initCoreServices(ctx context.Context, st *State, logger util.Logger) CoreServices {
	return CoreServices{
		User:   user.NewService(st.Files, logger),
		Script: scripting.NewService(st.Files, "scripts"),
		Socket: websocket.NewService(nil, nil),
	}
}
