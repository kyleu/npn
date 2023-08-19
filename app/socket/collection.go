package socket

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
	"github.com/pkg/errors"
	"strings"
)

func handleCollectionMessage(s *websocket.Service, c *websocket.Connection, cmd string, param json.RawMessage) error {
	switch cmd {
	case ClientMessageGetCollection:
		return getCollDetails(s, c, param)
	case ClientMessageAddCollection:
		return addCollection(s, c, param)
	case ClientMessageSaveCollection:
		return saveCollection(s, c, param)
	case ClientMessageDeleteCollection:
		return deleteCollection(s, c, param)
	case ClientMessageAddRequestURL:
		return addRequestURL(s, c, param)
	case ClientMessageTransform:
		return onTransformCollection(c, param, s)
	default:
		return errors.New("unhandled collection command [" + cmd + "]")
	}
}

func sendCollections(s *websocket.Service, c *websocket.Connection) {
	svcs := ctx(s)
	colls, err := svcs.Collection.Counts(&c.Profile.UserID)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error retrieving collections: %+v", err))
	}
	msg := websocket.NewMessage("collection", ServerMessageCollections, colls)
	err = s.WriteMessage(c.ID, msg)
	if err != nil {
		s.Logger.Warn(fmt.Sprintf("error writing to socket: %+v", err))
	}
}

func addCollection(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	name, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	if name == "" {
		name = "new"
	}
	key := util.Slugify(name)
	svcs := ctx(s)
	curr, _ := svcs.Collection.Load(&c.Profile.UserID, key)
	if curr != nil {
		key += "-" + strings.ToLower(util.RandomString(4))
	}

	err = svcs.Collection.Save(&c.Profile.UserID, "", key, name, "")
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+key+"]")
	}

	newColls, _ := svcs.Collection.Counts(&c.Profile.UserID)

	ret := &addCollOut{Collections: newColls, Active: key}
	msg := websocket.NewMessage("collection", ServerMessageCollectionAdded, ret)
	return s.WriteMessage(c.ID, msg)
}

func saveCollection(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	p := &saveCollOut{}
	err := util.FromJSONStrict(param, p)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}

	svcs := ctx(s)
	err = svcs.Collection.Save(&c.Profile.UserID, p.OriginalKey, p.Coll.Key, p.Coll.Title, p.Coll.Description)
	if err != nil {
		return errors.Wrap(err, "unable to save new collection with key ["+p.Coll.Key+"]")
	}

	msg := websocket.NewMessage("collection", ServerMessageCollectionUpdated, p.Coll)
	return s.WriteMessage(c.ID, msg)
}

func deleteCollection(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, "unable to parse input")
	}
	svcs := ctx(s)
	err = svcs.Collection.Delete(&c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, "unable to delete collection with key ["+key+"]")
	}

	msg := websocket.NewMessage("collection", ServerMessageCollectionDeleted, key)
	return s.WriteMessage(c.ID, msg)
}

func parseCollDetails(s *websocket.Service, userID *uuid.UUID, key string) (*collDetailsOut, error) {
	svcs := ctx(s)
	coll, err := svcs.Collection.Load(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	if coll == nil {
		return nil, nil
	}
	reqs, err := svcs.Request.List(userID, key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error retrieving requests for collection [%v]: %+v", key, err))
	}
	cd := &collDetailsOut{Key: key, Collection: coll, Requests: reqs}
	return cd, nil
}

func sendRequests(s *websocket.Service, c *websocket.Connection) {
	cd, _ := parseCollDetails(s, &c.Profile.UserID, "_")
	if cd == nil {
		msg := websocket.NewMessage("collection", ServerMessageCollectionNotFound, &cd)
		_ = s.WriteMessage(c.ID, msg)
	}
	msg := websocket.NewMessage("collection", ServerMessageCollectionDetail, cd)
	_ = s.WriteMessage(c.ID, msg)
}

func getCollDetails(s *websocket.Service, c *websocket.Connection, param json.RawMessage) error {
	key, err := util.FromJSONString(param)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error parsing collection [%v]: %+v", key, err))
	}
	cd, err := parseCollDetails(s, &c.Profile.UserID, key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error retrieving collection [%v]: %+v", key, err))
	}
	if cd == nil {
		msg := websocket.NewMessage("collection", ServerMessageCollectionNotFound, &cd)
		return s.WriteMessage(c.ID, msg)
	}
	msg := websocket.NewMessage("collection", ServerMessageCollectionDetail, cd)
	return s.WriteMessage(c.ID, msg)
}
