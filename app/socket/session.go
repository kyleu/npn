package socket

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/request/transform"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) handleSessionMessage(c *websocket.Connection, cmd string, param json.RawMessage, logger util.Logger) error {
	var err error

	switch cmd {
	case ClientMessageGetSession:
		err = s.onGetSession(c, param, logger)
	case ClientMessageAddSession:
		err = s.onAddSession(c, param, logger)
	case ClientMessageSaveSession:
		err = s.onSaveSession(c, param, logger)
	case ClientMessageDeleteSession:
		err = s.onDeleteSession(c, param, logger)
	case ClientMessageTransform:
		return s.onTransformSession(c, param, logger)
	default:
		err = errors.New("invalid session command [" + cmd + "]")
	}

	return err
}

func (s *Service) sendSessions(c *websocket.Connection, logger util.Logger) {
	sessions, err := s.Session.Counts(&c.Profile.ID)
	if err != nil {
		logger.Warn(fmt.Sprintf("error retrieving sessions: %+v", err))
	}
	msg := websocket.NewMessage(&c.Profile.ID, "session", ServerMessageSessions, sessions)
	err = s.Socket.WriteMessage(c.ID, msg, logger)
	if err != nil {
		logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

func (s *Service) onGetSession(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to read session key")
	}

	sess, err := s.Session.Load(&c.Profile.ID, key)
	if err != nil {
		return errors.Wrap(err, "unable to load session ["+key+"]")
	}
	var msg *websocket.Message
	if sess == nil {
		msg = websocket.NewMessage(&c.Profile.ID, "session", ServerMessageSessionNotFound, key)
	} else {
		msg = websocket.NewMessage(&c.Profile.ID, "session", ServerMessageSessionDetail, sess)
	}
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onAddSession(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	name, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse session name")
	}
	if name == "" {
		name = "new"
	}
	key := util.Slugify(name)
	curr, _ := s.Session.Load(&c.Profile.ID, key)
	if curr != nil {
		key += "-" + strings.ToLower(util.RandomString(4))
	}

	sess := &session.Session{Key: key, Title: name}
	err = s.Session.Save(&c.Profile.ID, "", sess)
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	sessCounts, _ := s.Session.Counts(&c.Profile.ID)

	ret := &addSessionOut{Sessions: sessCounts, Active: sess}
	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageSessionAdded, ret)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onSaveSession(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := &saveSessionIn{}
	err := util.FromJSONStrict(param, frm)
	if err != nil {
		return errors.Wrap(err, "can't load saveSession param")
	}
	frm.Sess = frm.Sess.Minify()
	err = s.Session.Save(&c.Profile.ID, frm.Orig, frm.Sess)
	if err != nil {
		return errors.Wrap(err, "can't save session ["+frm.Sess.Key+"]")
	}
	msg := websocket.NewMessage(&c.Profile.ID, "session", ServerMessageSessionDetail, frm.Sess.Normalize(frm.Sess.Key))
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onDeleteSession(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	err = s.Session.Delete(&c.Profile.ID, key)
	if err != nil {
		return errors.Wrap(err, "unable to delete session with key ["+key+"]")
	}

	msg := websocket.NewMessage(&c.Profile.ID, "session", ServerMessageSessionDeleted, key)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) onTransformSession(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	frm := ""
	err := util.FromJSONStrict(param, &frm)
	if err != nil {
		return errors.Wrap(err, "can't load session transform param")
	}

	sess, err := s.Session.Load(&c.Profile.ID, frm)
	if err != nil {
		return errors.Wrap(err, "can't load session transform ["+frm+"]")
	}

	rsp, err := transform.Session(sess, logger)
	if err != nil {
		return errors.Wrap(err, "can't load transform session")
	}

	msg := websocket.NewMessage(&c.Profile.ID, "session", ServerMessageSessionTransform, rsp)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}
