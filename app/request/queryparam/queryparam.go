package queryparam

import (
	"github.com/kyleu/npn/app/util"
	"net/url"
	"strings"
)

// Respresent an entry in a URL's querystring along with a description
type QueryParam struct {
	Key         string `json:"k,omitempty"`
	Value       string `json:"v,omitempty"`
	Description string `json:"desc,omitempty"`
}

// Returns a string in URL-encoded querystring format
func (q *QueryParam) String() string {
	return url.QueryEscape(q.Key) + "=" + url.QueryEscape(q.Value)
}

// Returns a cloned QueryParam with the provided data used as overrides
func (q *QueryParam) Merge(data util.ValueMap, logger util.Logger) *QueryParam {
	return &QueryParam{
		Key:         util.MergeLog("query."+q.Key+".key", q.Key, data, logger),
		Value:       util.MergeLog("query."+q.Key+".value", q.Value, data, logger),
		Description: util.MergeLog("query."+q.Key+".description", q.Description, data, logger),
	}
}

// Returns a shallow copy on this QueryParam
func (q *QueryParam) Clone() *QueryParam {
	return &QueryParam{Key: q.Key, Value: q.Value, Description: q.Description}
}

// Helper for arrays, represents a URL's querystring
type QueryParams []*QueryParam

// Returns a string in URL-encoded querystring format
func (q QueryParams) String() string {
	ret := make([]string, 0, len(q))
	for _, x := range q {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, "&")
}

// Returns a cloned QueryParams with the provided data used as overrides
func (q QueryParams) Merge(data util.ValueMap, logger util.Logger) QueryParams {
	ret := make(QueryParams, 0, len(q))
	for _, qp := range q {
		ret = append(ret, qp.Merge(data, logger))
	}
	return ret
}

// Returns a shallow copy on this QueryParams
func (q QueryParams) Clone() QueryParams {
	ret := make(QueryParams, 0, len(q))
	for _, qp := range q {
		ret = append(ret, qp.Clone())
	}
	return ret
}

// Parses the provided URL querystring
func QueryParamsFromRaw(s string) QueryParams {
	ret := make(QueryParams, 0)
	parts := strings.Split(s, "&")
	for _, x := range parts {
		if len(x) > 0 {
			k, v := util.StringSplit(x, '=', true)
			k, _ = url.QueryUnescape(k)
			v, _ = url.QueryUnescape(v)
			ret = append(ret, &QueryParam{Key: k, Value: v})
		}
	}
	return ret
}
