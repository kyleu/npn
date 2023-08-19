package socket

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request/search"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) handleSystemMessage(c *websocket.Connection, cmd string, param json.RawMessage, logger util.Logger) error {
	switch cmd {
	case ClientMessageSearch:
		return s.onSearch(c, param, logger)
	case ClientMessageSaveProfile:
		return s.saveProfile(c, param)
	default:
		return errors.New("unhandled system command [" + cmd + "]")
	}
}

func (s *Service) onSearch(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	sp := &search.Params{}
	err := util.FromJSON(param, sp)
	if err != nil {
		return errors.Wrap(err, "unable to parse search params")
	}
	results, err := s.Search.Run(sp, &c.Profile.ID)
	if err != nil {
		return errors.Wrap(err, "search error")
	}
	msg := websocket.NewMessage(&c.Profile.ID, "system", ServerMessageSearchResults, results)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) saveProfile(c *websocket.Connection, param json.RawMessage) error {
	return nil
}
