package request

import (
	"github.com/kyleu/npn/app/util"
	"github.com/pkg/errors"
	"strings"
)

type Request struct {
	Key         string     `json:"key,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Prototype   *Prototype `json:"prototype"`
}

func FromString(key string, content string) (*Request, error) {
	ret := &Request{}
	content = strings.TrimSpace(content)
	if strings.HasPrefix(content, "{") {
		b := []byte(content)
		errRequest := util.FromJSON(b, ret)
		if errRequest != nil {
			proto := &Prototype{}
			errProto := util.FromJSON(b, proto)
			if errProto != nil {
				return nil, errors.Wrapf(errRequest, "unable to parse request from [%s]", content)
			}
			ret.Prototype = proto
		}
	} else {
		u := strings.TrimPrefix(strings.TrimSuffix(content, `"`), `"`)
		proto := PrototypeFromString(u)
		ret.Prototype = proto
	}
	return ret.Normalize(key), nil
}

func (r *Request) TitleWithFallback() string {
	if r.Title == "" {
		return r.Key
	}
	return r.Title
}

func (r *Request) Options() *Options {
	if r.Prototype == nil {
		return &Options{}
	}
	if r.Prototype.Options == nil {
		return &Options{}
	}
	return r.Prototype.Options
}

func (r *Request) Merge(data util.ValueMap, logger util.Logger) *Request {
	return &Request{
		Key:         r.Key,
		Title:       util.MergeLog("title", r.Title, data, logger),
		Description: util.MergeLog("description", r.Description, data, logger),
		Prototype:   r.Prototype.Merge(data, logger),
	}
}

type Requests []*Request
