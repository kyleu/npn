package socket

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/call"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/imprt"
	"github.com/kyleu/npn/app/request/search"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/user"
	"github.com/kyleu/npn/app/util"
)

type Service struct {
	User       *user.Service
	Session    *session.Service
	Collection *collection.Service
	Request    *request.Service
	Caller     *call.Service
	Search     *search.Service
	Import     *imprt.Service
	Socket     *websocket.Service
}

func NewService(
	us *user.Service, ss *session.Service, co *collection.Service, rq *request.Service, cl *call.Service, sr *search.Service, im *imprt.Service,
) *Service {
	ret := &Service{User: us, Session: ss, Collection: co, Request: rq, Caller: cl, Search: sr, Import: im}
	ret.Socket = websocket.NewService(ret.onOpen, nil)
	return ret
}

func (s *Service) Handler(
	ctx context.Context, ws *websocket.Service, c *websocket.Connection, svc string, cmd string, param []byte, logger util.Logger,
) error {
	var err error
	switch svc {
	case "system":
		err = s.handleSystemMessage(c, cmd, param, logger)
	case "collection":
		err = s.handleCollectionMessage(c, cmd, param, logger)
	case "request":
		err = s.handleRequestMessage(c, cmd, param, logger)
	case "session":
		err = s.handleSessionMessage(c, cmd, param, logger)
	case "import":
		err = s.handleImportMessage(c, cmd, param, logger)
	default:
		err = errors.Errorf("invalid service ID [%s]", svc)
	}
	return errors.Wrap(err, "error handling socket message ["+cmd+"]")
}

func (s *Service) onOpen(_ *websocket.Service, c *websocket.Connection, logger util.Logger) error {
	go s.sendCollections(c, logger)
	go s.sendRequests(c, logger)
	go s.sendSessions(c, logger)
	return nil
}

func onClose(*websocket.Service, *websocket.Connection) error {
	return nil
}
