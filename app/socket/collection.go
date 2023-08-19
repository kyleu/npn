package socket

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) handleCollectionMessage(c *websocket.Connection, cmd string, param json.RawMessage, logger util.Logger) error {
	switch cmd {
	case ClientMessageGetCollection:
		return s.getCollDetails(c, param, logger)
	case ClientMessageAddCollection:
		return s.addCollection(c, param, logger)
	case ClientMessageSaveCollection:
		return s.saveCollection(c, param, logger)
	case ClientMessageDeleteCollection:
		return s.deleteCollection(c, param, logger)
	case ClientMessageAddRequestURL:
		return s.addRequestURL(c, param, logger)
	case ClientMessageTransform:
		return s.onTransformCollection(c, param, logger)
	default:
		return errors.New("unhandled collection command [" + cmd + "]")
	}
}

func (s *Service) sendCollections(c *websocket.Connection, logger util.Logger) {
	colls, err := s.Collection.Counts(&c.Profile.ID)
	if err != nil {
		logger.Warn(fmt.Sprintf("error retrieving collections: %+v", err))
	}
	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollections, colls)
	err = s.Socket.WriteMessage(c.ID, msg, logger)
	if err != nil {
		logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

func (s *Service) addCollection(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	name, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	if name == "" {
		name = "new"
	}
	key := util.Slugify(name)
	curr, _ := s.Collection.Load(&c.Profile.ID, key)
	if curr != nil {
		key += "-" + strings.ToLower(util.RandomString(4))
	}

	err = s.Collection.Save(&c.Profile.ID, "", key, name, "")
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	newColls, _ := s.Collection.Counts(&c.Profile.ID)

	ret := &addCollOut{Collections: newColls, Active: key}
	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionAdded, ret)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) saveCollection(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	p := &saveCollOut{}
	err := util.FromJSONStrict(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}

	err = s.Collection.Save(&c.Profile.ID, p.OriginalKey, p.Coll.Key, p.Coll.Title, p.Coll.Description)
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+p.Coll.Key+"]")
	}

	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionUpdated, p.Coll)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) deleteCollection(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	err = s.Collection.Delete(&c.Profile.ID, key)
	if err != nil {
		return errors.Wrap(err, "unable to delete collection with key ["+key+"]")
	}

	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionDeleted, key)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) parseCollDetails(userID *uuid.UUID, key string) (*collDetailsOut, error) {
	coll, err := s.Collection.Load(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	if coll == nil {
		return nil, nil
	}
	reqs, err := s.Request.List(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving requests for collection [%v]: %+v", key, err))
	}
	cd := &collDetailsOut{Key: key, Collection: coll, Requests: reqs}
	return cd, nil
}

func (s *Service) sendRequests(c *websocket.Connection, logger util.Logger) {
	cd, _ := s.parseCollDetails(&c.Profile.ID, "_")
	if cd == nil {
		msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionNotFound, &cd)
		_ = s.Socket.WriteMessage(c.ID, msg, logger)
	}
	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionDetail, cd)
	_ = s.Socket.WriteMessage(c.ID, msg, logger)
}

func (s *Service) getCollDetails(c *websocket.Connection, param json.RawMessage, logger util.Logger) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing collection [%v]: %+v", key, err))
	}
	cd, err := s.parseCollDetails(&c.Profile.ID, key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	if cd == nil {
		msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionNotFound, &cd)
		return s.Socket.WriteMessage(c.ID, msg, logger)
	}
	msg := websocket.NewMessage(&c.Profile.ID, "collection", ServerMessageCollectionDetail, cd)
	return s.Socket.WriteMessage(c.ID, msg, logger)
}
