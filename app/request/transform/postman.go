package transform

import (
	"fmt"
	"strings"

	"github.com/rbretecher/go-postman-collection"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/util"
)

type Postman struct{}

var _ RequestTransformer = (*Postman)(nil)

func (x *Postman) Key() string {
	return "postman"
}

func (x *Postman) Title() string {
	return "Postman"
}

func (x *Postman) Description() string {
	return "TODO: postman"
}

func (x *Postman) TransformRequest(proto *request.Prototype, sess *session.Session, logger util.Logger) (*Result, error) {
	pReq := toPostmanRequest(proto)
	out := util.ToJSON(pReq)
	return &Result{Out: out}, nil
}

func (x *Postman) TransformCollection(c *collection.FullCollection, logger util.Logger) (*Result, error) {
	pColl := postman.CreateCollection(c.Coll.TitleWithFallback(), c.Coll.Description)
	for _, r := range c.Requests {
		pColl.AddItem(postman.CreateItem(postman.Item{
			Name:        r.TitleWithFallback(),
			Description: r.Description,
			ID:          r.Key,
			Request:     toPostmanRequest(r.Prototype),
			Responses:   nil,
		}))
	}
	out := util.ToJSON(pColl)
	return &Result{Out: out}, nil
}

func toPostmanURL(p *request.Prototype) *postman.URL {
	return &postman.URL{
		Raw:       p.URLString(),
		Protocol:  p.Protocol.String(),
		Host:      []string{p.Domain},
		Path:      strings.Split(p.Path, "/"),
		Port:      fmt.Sprintf("%v", p.Port),
		Query:     nil,
		Hash:      p.Fragment,
		Variables: nil,
	}
}

func toPostmanRequest(p *request.Prototype) *postman.Request {
	return &postman.Request{
		URL:         toPostmanURL(p),
		Auth:        nil,
		Proxy:       nil,
		Certificate: nil,
		// Method:      ???,
		Description: nil,
		Header:      nil,
		Body:        nil,
	}
}

var txPostman = &Postman{}
