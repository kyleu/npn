package body

import "github.com/kyleu/npn/app/util"

const KeyJSON = "json"

type JSON struct {
	Msg any `json:"msg"`
	str string
}

var _ Config = (*JSON)(nil)

func NewJSON(msg any) *Body {
	return NewBody(KeyJSON, &JSON{Msg: msg})
}

func parseJSON(ct string, charset string, b []byte) *Body {
	x, err := util.FromJSONAny(b)
	if err != nil {
		if ct == "" {
			return NewError(err.Error())
		}
		return detect("", "", charset, b)
	}

	return NewJSON(x)
}

func (j *JSON) ContentLength() int64 {
	return int64(len(j.Bytes()))
}

func (j *JSON) Bytes() []byte {
	return []byte(j.String())
}

func (j *JSON) MimeType() string {
	return "application/json"
}

func (j *JSON) String() string {
	if j.str == "" {
		j.str = util.ToJSON(j.Msg)
	}
	return j.str
}

func (j *JSON) Merge(data util.ValueMap, logger util.Logger) Config {
	m := j.Msg
	ms := util.ToJSONCompact(j.Msg)
	if len(ms) > 0 {
		ms = util.MergeLog("body.json.msg", ms, data, logger)
		i, err := util.FromJSONAny([]byte(ms))
		if err == nil && i != nil {
			m = i
		}
	}

	s := j.str
	if len(s) > 0 {
		s = util.MergeLog("body.json.str", s, data, logger)
	}
	return &JSON{Msg: m, str: s}
}

func (j *JSON) Clone() *Body {
	return NewBody(KeyJSON, &JSON{Msg: j.Msg, str: j.str})
}
