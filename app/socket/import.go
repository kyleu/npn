package socket

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) handleImportMessage(c *websocket.Connection, cmd string, param json.RawMessage, logger util.Logger) error {
	switch cmd {
	case ClientMessageGetImport:
		return s.getImport(c, param, logger)
	default:
		return errors.New("unhandled import command [" + cmd + "]")
	}
}

func (s *Service) getImport(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse import key")
	}
	cfg, results, err := s.Import.Load(key)
	if err != nil {
		return errors.Wrap(err, "can't load import ["+key+"]")
	}
	ret := map[string]any{"key": key, "cfg": cfg, "results": results}
	msg := websocket.NewMessage(&c.Profile.ID, "import", ServerMessageImportResult, ret)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}
