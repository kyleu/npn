package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/imprt"
	"github.com/kyleu/npn/app/request/session"
	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/request"

	"github.com/kyleu/npn/app/call"
)

type Dependencies struct {
	User       user.Service
	Session    *session.Service
	Collection *collection.Service
	Request    *request.Service
	Caller     *call.Service
	Search     *search.Service
	Import     *imprt.Service
}

func NewService(deps *Dependencies, logger *logrus.Logger) *websocket.Service {
	return websocket.NewService(logger, onOpen, handler, onClose, deps)
}

func handler(s *websocket.Service, c *websocket.Connection, svc string, cmd string, param json.RawMessage) error {
	var err error
	switch svc {
	case "system":
		err = handleSystemMessage(s, c, cmd, param)
	case "collection":
		err = handleCollectionMessage(s, c, cmd, param)
	case "request":
		err = handleRequestMessage(s, c, cmd, param)
	case "session":
		err = handleSessionMessage(s, c, cmd, param)
	case "import":
		err = handleImportMessage(s, c, cmd, param)
	default:
		err = errors.Errorf("invalid service ID [%s]", svc)
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

func onOpen(s *websocket.Service, c *websocket.Connection) error {
	go sendCollections(s, c)
	go sendRequests(s, c)
	go sendSessions(s, c)
	return nil
}

func onClose(*websocket.Service, *websocket.Connection) error {
	return nil
}

func ctx(s *websocket.Service) *Dependencies {
	return s.Context.(*Dependencies)
}
