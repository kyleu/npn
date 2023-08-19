package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
	"github.com/pkg/errors"
)

func handleImportMessage(s *websocket.Service, c *websocket.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetImport:
		return getImport(s, c, param)
	default:
		return errors.New("unhandled import command [" + cmd + "]")
	}
}

func getImport(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse import key")
	}
	cfg, results, err := ctx(s).Import.Load(key)
	if err != nil {
		return errors.Wrap(err, "can't load import ["+key+"]")
	}
	ret := map[string]interface{}{"key": key, "cfg": cfg, "results": results}
	msg := websocket.NewMessage("import", ServerMessageImportResult, ret)
	return s.WriteMessage(c.ID, msg)
}
