package socket

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
)

func handleRequestMessage(s *websocket.Service, c *websocket.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageRunURL:
		return onRunURL(c, param, s)
	case ClientMessageGetRequest:
		return onGetRequest(c, param, s)
	case ClientMessageSaveRequest:
		return onSaveRequest(c, param, s)
	case ClientMessageDeleteRequest:
		return onDeleteRequest(c, param, s)
	case ClientMessageCall:
		return onCall(c, param, s)
	case ClientMessageTransform:
		return onTransformRequest(c, param, s)
	default:
		return errors.New("invalid request command [" + cmd + "]")
	}
}

func onRunURL(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	url, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to read URL")
	}
	return AddRequestFromURL(s, c, "_", url)
}

func onGetRequest(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	svc := ctx(s)
	frm := &getRequestIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load getRequest param")
	}
	req, err := svc.Request.Load(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		msg := websocket.NewMessage("request", ServerMessageRequestNotFound, frm)
		return s.WriteMessage(c.ID, msg)
	}
	ret := &reqDetailOut{Coll: frm.Coll, OrigKey: req.Key, Req: req}
	msg := websocket.NewMessage("request", ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onSaveRequest(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	svc := ctx(s)
	frm := &saveRequestIn{}
	err := util.FromJSON(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	frm.Req = frm.Req.Minify()
	err = svc.Request.Save(&c.Profile.UserID, frm.Coll, frm.Orig, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't save request ["+frm.Req.Key+"]")
	}
	ret := &reqDetailOut{Coll: frm.Coll, OrigKey: frm.Orig, Req: frm.Req}
	msg := websocket.NewMessage("request", ServerMessageRequestDetail, ret)
	return s.WriteMessage(c.ID, msg)
}

func onDeleteRequest(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	svc := ctx(s)
	frm := &deleteRequestIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveRequest param")
	}
	err = svc.Request.Delete(&c.Profile.UserID, frm.Coll, frm.Req)
	if err != nil {
		return errors.Wrap(err, "can't remove request")
	}

	summaries, err := svc.Request.List(&c.Profile.UserID, frm.Coll)
	if err != nil {
		return errors.Wrap(err, "can't list requests")
	}

	ret := &reqDeletedOut{Coll: frm.Coll, Req: frm.Req, Requests: summaries}
	msg := websocket.NewMessage("request", ServerMessageRequestDeleted, ret)
	return s.WriteMessage(c.ID, msg)
}

func onCall(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	svc := ctx(s)
	frm := &callIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load request call param")
	}

	sess, err := ctx(s).Session.Load(&c.Profile.UserID, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't load session ["+frm.Sess+"]")
	}

	go func() {
		onStarted := func(started *call.RequestStarted) {
			msg := websocket.NewMessage("request", ServerMessageRequestStarted, started)
			_ = s.WriteMessage(c.ID, msg)
		}
		onCompleted := func(completed *call.RequestCompleted) {
			msg := websocket.NewMessage("request", ServerMessageRequestCompleted, completed)
			_ = s.WriteMessage(c.ID, msg)
		}
		err := svc.Caller.Call(&c.Profile.UserID, frm.Coll, frm.Req, frm.Proto, sess, onStarted, onCompleted)
		if err != nil {
			s.Logger.Warn(fmt.Sprintf("error calling [%v]: %+v", frm.Req, err))
		}
	}()

	return nil
}
