package socket

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/search"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
)

func handleSystemMessage(s *websocket.Service, c *websocket.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageTestbed:
		return testbed(param, s)
	case ClientMessageSearch:
		return onSearch(s, c, param)
	case ClientMessageSaveProfile:
		return saveProfile(s, c, param)
	default:
		return errors.New("unhandled system command [" + cmd + "]")
	}
}

func onSearch(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	sp := &search.Params{}
	err := util.FromJSON(param, sp)
	if err != nil {
		return errors.Wrap(err, "unable to parse search params")
	}
	results, err := ctx(s).Search.Run(sp, &c.Profile.UserID, c.Profile.Role)
	if err != nil {
		return errors.Wrap(err, "search error")
	}
	msg := websocket.NewMessage("system", ServerMessageSearchResults, results)
	return s.WriteMessage(c.ID, msg)
}

func saveProfile(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	p := &npnuser.Profile{}
	err := util.FromJSON(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse profile")
	}
	p.UserID = c.Profile.UserID
	p.Role = c.Profile.Role

	svcs := ctx(s)

	_, err = svcs.User.SaveProfile(p.ToUserProfile())
	return err
}
