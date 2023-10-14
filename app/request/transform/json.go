package transform

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/util"
)

type JSON struct{}

var _ RequestTransformer = (*JSON)(nil)

func (x *JSON) Key() string {
	return "json"
}

func (x *JSON) Title() string {
	return "JSON"
}

func (x *JSON) Description() string {
	return "TODO: json"
}

func (x *JSON) TransformRequest(proto *request.Prototype, sess *session.Session, logger util.Logger) (*Result, error) {
	out := util.ToJSON(proto)
	return &Result{Out: out}, nil
}

func (x *JSON) TransformCollection(c *collection.FullCollection, logger util.Logger) (*Result, error) {
	src := map[string]any{"coll": c.Coll, "requests": c.Requests}
	out := util.ToJSON(src)
	return &Result{Out: out}, nil
}

var txJSON = &JSON{}
