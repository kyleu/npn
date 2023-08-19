package socket

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) handleRequestMessage(c *websocket.Connection, cmd string, param json.RawMessage, logger util.Logger) error {
	switch cmd {
	case ClientMessageRunURL:
		return s.onRunURL(c, param, logger)
	case ClientMessageGetRequest:
		return s.onGetRequest(c, param, logger)
	case ClientMessageSaveRequest:
		return s.onSaveRequest(c, param, logger)
	case ClientMessageDeleteRequest:
		return s.onDeleteRequest(c, param, logger)
	case ClientMessageCall:
		return s.onCall(c, param, logger)
	case ClientMessageTransform:
		return s.onTransformRequest(c, param, logger)
	default:
		return errors.New("invalid request command [" + cmd + "]")
	}
}

func (s *Service) onRunURL(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	url, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to read URL")
	}
	return s.AddRequestFromURL(c, "_", url, logger)
}

func (s *Service) onGetRequest(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &getRequestIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load getRequest param")
	}
	req, err := s.Request.Load(&c.Profile.ID, frm.Coll, frm.Req)
	if err != nil {
		msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestNotFound, frm)
		return s.Socket.WriteMessage(c.ID, msg, logger)
	}
	ret := &reqDetailOut{Coll: frm.Coll, OrigKey: req.Key, Req: req}
	msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestDetail, ret)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onSaveRequest(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &saveRequestIn{}
	err := util.FromJSON(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	frm.Req = frm.Req.Minify()
	err = s.Request.Save(&c.Profile.ID, frm.Coll, frm.Orig, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't save request ["+frm.Req.Key+"]")
	}
	ret := &reqDetailOut{Coll: frm.Coll, OrigKey: frm.Orig, Req: frm.Req}
	msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestDetail, ret)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onDeleteRequest(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &deleteRequestIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = s.Request.Delete(&c.Profile.ID, frm.Coll, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't remove request")
	}

	summaries, err := s.Request.List(&c.Profile.ID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't list requests")
	}

	ret := &reqDeletedOut{Coll: frm.Coll, Req: frm.Req, Requests: summaries}
	msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestDeleted, ret)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onCall(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &callIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request call param")
	}

	sess, err := s.Session.Load(&c.Profile.ID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load session ["+frm.Sess+"]")
	}

	go func() {
		onStarted := func(started *call.RequestStarted) {
			msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestStarted, started)
			_ = s.Socket.WriteMessage(c.ID, msg, logger)
		}
		onCompleted := func(completed *call.RequestCompleted) {
			msg := websocket.NewMessage(&c.Profile.ID, "request", ServerMessageRequestCompleted, completed)
			_ = s.Socket.WriteMessage(c.ID, msg, logger)
		}
		err := s.Caller.Call(&c.Profile.ID, frm.Coll, frm.Req, frm.Proto, sess, onStarted, onCompleted)
		if err != nil {
			logger.Warn(fmt.Sprintf("error calling [%v]: %+v", frm.Req, err))
		}
	}()

	return nil
}
