package socket

import (
	"encoding/json"
	"fmt"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/request/transform"
	"github.com/kyleu/npn/app/util"
	"github.com/pkg/errors"
	"strings"
)

func handleSessionMessage(s *websocket.Service, c *websocket.Connection, cmd string, param json.RawMessage) error {
	var err error

	switch cmd {
	case ClientMessageGetSession:
		err = onGetSession(c, param, s)
	case ClientMessageAddSession:
		err = onAddSession(c, param, s)
	case ClientMessageSaveSession:
		err = onSaveSession(c, param, s)
	case ClientMessageDeleteSession:
		err = onDeleteSession(c, param, s)
	case ClientMessageTransform:
		return onTransformSession(c, param, s)
	default:
		err = errors.New("invalid session command [" + cmd + "]")
	}

	return err
}

func sendSessions(s *websocket.Service, c *websocket.Connection) {
	svcs := ctx(s)
	sessions, err := svcs.Session.Counts(&c.Profile.UserID)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error retrieving sessions: %+v", err))
	}
	msg := websocket.NewMessage("session", ServerMessageSessions, sessions)
	err = s.WriteMessage(c.ID, msg)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

func onGetSession(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to read session key")
	}

	svcs := ctx(s)
	sess, err := svcs.Session.Load(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to load session ["+key+"]")
	}
	var msg *websocket.Message
	if sess == nil {
		msg = websocket.NewMessage("session", ServerMessageSessionNotFound, key)
	} else {
		msg = websocket.NewMessage("session", ServerMessageSessionDetail, sess)
	}
	return s.WriteMessage(c.ID, msg)
}

func onAddSession(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	name, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse session name")
	}
	if name == "" {
		name = "new"
	}
	key := util.Slugify(name)
	svcs := ctx(s)
	curr, _ := svcs.Session.Load(&c.Profile.UserID, key)
	if curr != nil {
		key += "-" + strings.ToLower(util.RandomString(4))
	}

	sess := &session.Session{Key: key, Title: name}
	err = svcs.Session.Save(&c.Profile.UserID, "", sess)
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	sessCounts, _ := svcs.Session.Counts(&c.Profile.UserID)

	ret := &addSessionOut{Sessions: sessCounts, Active: sess}
	msg := websocket.NewMessage("collection", ServerMessageSessionAdded, ret)
	return s.WriteMessage(c.ID, msg)
}

func onSaveSession(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	svc := ctx(s)
	frm := &saveSessionIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveSession param")
	}
	frm.Sess = frm.Sess.Minify()
	err = svc.Session.Save(&c.Profile.UserID, frm.Orig, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't save session ["+frm.Sess.Key+"]")
	}
	msg := websocket.NewMessage("session", ServerMessageSessionDetail, frm.Sess.Normalize(frm.Sess.Key))
	return s.WriteMessage(c.ID, msg)
}

func onDeleteSession(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	svcs := ctx(s)
	err = svcs.Session.Delete(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to delete session with key ["+key+"]")
	}

	msg := websocket.NewMessage("session", ServerMessageSessionDeleted, key)
	return s.WriteMessage(c.ID, msg)
}

func onTransformSession(c *websocket.Connection, param json.RawMessage, s *websocket.Service) error {
	frm := ""
	err := util.FromJSONStrict(param, &frm)
	if err != nil {
		return errors.Wrap(err, "can't load session transform param")
	}

	sess, err := ctx(s).Session.Load(&c.Profile.UserID, frm)
	if err != nil {
		return errors.Wrap(err, "can't load session transform ["+frm+"]")
	}

	rsp, err := transform.Session(sess, s.Logger)
	if err != nil {
		return errors.Wrap(err, "can't load transform session")
	}

	msg := websocket.NewMessage("session", ServerMessageSessionTransform, rsp)
	return s.WriteMessage(c.ID, msg)
}
