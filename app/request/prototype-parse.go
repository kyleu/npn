package request

import (
	"github.com/kyleu/npn/app/request/auth"
	"github.com/kyleu/npn/app/request/queryparam"
	"github.com/kyleu/npn/app/util"
	"net/url"
	"strconv"
	"strings"
)

func PrototypeFromURL(u *url.URL) *Prototype {
	var at *auth.Auth
	if u.User != nil {
		p, _ := u.User.Password()
		a := auth.NewBasic(u.User.Username(), p, false)
		at = a
	}
	domain, portString := util.StringSplit(u.Host, ':', true)

	port := 0
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(u.Scheme),
		Domain:   domain,
		Port:     port,
		Path:     u.Path,
		Query:    queryparam.QueryParamsFromRaw(u.RawQuery),
		Fragment: u.Fragment,
		Auth:     at,
	}
}

func PrototypeFromString(u string) *Prototype {
	var at *auth.Auth

	rest, frag := util.StringSplit(u, '#', true)
	if len(frag) > 0 {
		frag, _ = url.QueryUnescape(frag)
	}
	rest, query := util.StringSplit(rest, '?', true)
	proto, rest := util.StringSplit(rest, ':', true)
	if rest == "" {
		rest = proto
		proto = "http"
	}
	rest = strings.TrimPrefix(strings.TrimPrefix(rest, "/"), "/")
	rest, path := util.StringSplit(rest, '/', true)
	if len(path) > 0 {
		path, _ = url.PathUnescape(path)
	}
	aut, host := util.StringSplit(rest, '@', true)
	if host == "" {
		host = aut
		aut = ""
	}
	host, portString := util.StringSplit(host, ':', true)
	port := 0
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	if aut != "" {
		user, pass := util.StringSplit(aut, ':', true)
		a := auth.NewBasic(user, pass, false)
		at = a
	}
	return &Prototype{
		Method:   MethodGet,
		Protocol: ProtocolFromString(proto),
		Domain:   host,
		Port:     port,
		Path:     path,
		Query:    queryparam.QueryParamsFromRaw(query),
		Fragment: frag,
		Auth:     at,
	}
}
