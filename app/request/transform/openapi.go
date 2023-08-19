package transform

import (
	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/util"
)

type OpenAPI struct {
	Multiline bool
}

var _ CollectionTransformer = (*OpenAPI)(nil)

func (x *OpenAPI) Key() string {
	return "openapi"
}

func (x *OpenAPI) Title() string {
	return "OpenAPI"
}

func (x *OpenAPI) Description() string {
	return "TODO: OpenAPI"
}

func (x *OpenAPI) ApplyToMultiple() bool {
	return true
}

func (x *OpenAPI) TransformRequest(p *request.Prototype, sess *session.Session, logger util.Logger) (*Result, error) {
	out := "OpenAPI"
	return &Result{Out: out}, nil
}

func (x *OpenAPI) TransformCollection(f *collection.FullCollection, logger util.Logger) (*Result, error) {
	out := "OpenAPI: TODO!"
	return &Result{Out: out}, nil
}

var txOpenAPI = &OpenAPI{}
