package socket

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/transform"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) onTransformRequest(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &transformIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request transform param")
	}

	tx := transform.AllRequestTransformers.Get(frm.Fmt)
	if tx == nil {
		return errors.New("can't load request transformer [" + frm.Fmt + "]")
	}

	sess, err := s.Session.Load(&c.Profile.ID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load request transform session ["+frm.Sess+"]")
	}

	rsp, err := tx.TransformRequest(frm.Proto, sess, logger)
	if err != nil {
		return errors.Wrap(err, "can't transform request")
	}

	txr := transformOut{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestTransform, txr)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onTransformCollection(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &transformIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform param")
	}

	tx := transform.AllCollectionTransformers.Get(frm.Fmt)
	if tx == nil {
		return errors.New("can't load collection transformer [" + frm.Fmt + "]")
	}

	coll, err := s.Collection.Load(&c.Profile.ID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform collection ["+frm.Coll+"]")
	}

	requests, err := s.Request.LoadAll(&c.Profile.ID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform requests for ["+frm.Coll+"]")
	}

	sess, err := s.Session.Load(&c.Profile.ID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load collection transform session ["+frm.Sess+"]")
	}

	tc := &collection.FullCollection{Coll: coll, Requests: requests, Sess: sess}
	rsp, err := tx.TransformCollection(tc, logger)
	if err != nil {
		return errors.Wrap(err, "can't load transform colllection")
	}

	txr := transformOut{Coll: frm.Coll, Req: frm.Req, Fmt: frm.Fmt, Out: rsp.Out}
	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionTransform, txr)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}
