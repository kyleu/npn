package socket

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) addRequestURL(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	p := &addURLIn{}
	err := util.FromJSONStrict(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse input from URL")
	}
	return s.AddRequestFromURL(c, p.Coll, p.URL, logger)
}

func (s *Service) AddRequestFromURL(c *websocket.Connection, coll string, url string, logger util.Logger) error {
	req, err := request.FromString("new", url)
	if err != nil {
		return errors.Wrap(err, "unable to parse request from URL ["+url+"]")
	}
	req.Key = util.Slugify(req.Prototype.Domain)

	curr, _ := s.Request.Load(&c.Profile.ID, coll, req.Key)
	if curr != nil {
		clean(req)
		curr, _ = s.Request.Load(&c.Profile.ID, coll, req.Key)
		if curr != nil {
			req.Key += "-" + strings.ToLower(util.RandomString(4))
		}
	}

	err = s.Request.Save(&c.Profile.ID, coll, "", req)
	if err != nil {
		return errors.Wrap(err, "unable to save request from URL ["+url+"]")
	}

	x, err := s.parseCollDetails(&c.Profile.ID, coll)
	if err != nil {
		return err
	}

	out := &addURLOut{
		Coll: x,
		Req:  req,
	}
	msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestAdded, out)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func clean(req *request.Request) {
	if req.Title == "" {
		req.Title = req.Key
	}
	if req.Prototype != nil && len(req.Prototype.Path) > 0 {
		add := req.Prototype.Path
		if len(add) > 8 {
			add = add[0:8]
		}
		req.Key += "-" + util.Slugify(add)
	}
}
