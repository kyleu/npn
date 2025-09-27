package transform

import (
	"path"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/request"
	"github.com/kyleu/npn/app/request/collection"
	"github.com/kyleu/npn/app/request/session"
	"github.com/kyleu/npn/app/util"
)

func OpenAPI2Import(data []byte) (*openapi3.T, error) {
	x := &openapi2.T{}
	err := yaml.Unmarshal(data, x)
	if err != nil {
		return nil, err
	}
	swag, err := openapi2conv.ToV3(x)
	return swag, err
}

func OpenAPI3Import(data []byte) (*openapi3.T, error) {
	swag, err := openapi3.NewLoader().LoadFromData(data)
	return swag, err
}

func OpenAPIToFullCollection(swag *openapi3.T) (*collection.FullCollection, error) {
	coll := openAPIToCollection(swag.Info)
	reqs, sess, err := openAPIToRequests(swag)
	if err != nil {
		return nil, err
	}

	ret := &collection.FullCollection{Coll: coll, Requests: reqs, Sess: sess}
	return ret, nil
}

func openAPIToRequests(swag *openapi3.T) (request.Requests, *session.Session, error) {
	reqs := make(request.Requests, 0)
	sess := &session.Session{}

	url := ""
	if len(swag.Servers) > 0 {
		url = swag.Servers[0].URL
	}

	proto := request.PrototypeFromString(url)

	for k, p := range swag.Paths.Map() {
		newReqs, err := openAPIPathToRequests(k, p, proto)
		if err != nil {
			return nil, nil, err
		}
		reqs = append(reqs, newReqs...)
	}

	return reqs, sess, nil
}

func openAPIPathToRequests(pathKey string, pathItem *openapi3.PathItem, proto *request.Prototype) (request.Requests, error) {
	ret := make(request.Requests, 0)
	ops := pathItem.Operations()
	for meth, op := range ops {
		p := proto.Clone()
		p.Method = request.MethodFromString(meth)
		p.Path = path.Join(proto.Path, pathKey)

		rk := op.OperationID
		if rk == "" {
			rk = util.Slugify(op.Description)
		}
		if rk == "" {
			rk = util.Slugify(p.Method.Key + "_" + p.Path)
		}
		if rk == "" {
			return nil, errors.New("unable to determine action name")
		}

		ret = append(ret, &request.Request{
			Key:         rk,
			Title:       "",
			Description: "",
			Prototype:   p,
		})
	}
	return ret, nil
}

func openAPIToCollection(i *openapi3.Info) *collection.Collection {
	return &collection.Collection{Key: util.Slugify(i.Title), Title: i.Title, Description: i.Description}
}
