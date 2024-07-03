// Code generated by qtc from "URL.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/URL.html:1
package view

//line views/components/view/URL.html:1
import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
)

//line views/components/view/URL.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/URL.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/URL.html:11
func StreamURL(qw422016 *qt422016.Writer, u any, content string, includeExternalIcon bool, ps *cutil.PageState) {
//line views/components/view/URL.html:12
	if u == nil {
//line views/components/view/URL.html:12
		qw422016.N().S(`<em>nil</em>`)
//line views/components/view/URL.html:14
	} else {
//line views/components/view/URL.html:16
		var href string
		switch t := u.(type) {
		case string:
			href = t
		case url.URL:
			href = t.String()
		case *url.URL:
			href = t.String()
		default:
			href = fmt.Sprint(u)
		}
		if content == "" {
			content = href
		}
		showIcon := includeExternalIcon && strings.HasPrefix(href, "http")

//line views/components/view/URL.html:31
		qw422016.N().S(`<a target="_blank" rel="noopener noreferrer" href="`)
//line views/components/view/URL.html:32
		qw422016.E().S(href)
//line views/components/view/URL.html:32
		qw422016.N().S(`">`)
//line views/components/view/URL.html:32
		qw422016.E().S(content)
//line views/components/view/URL.html:32
		if showIcon {
//line views/components/view/URL.html:32
			components.StreamSVGLinkPadded(qw422016, `external`, ps)
//line views/components/view/URL.html:32
		}
//line views/components/view/URL.html:32
		qw422016.N().S(`</a>`)
//line views/components/view/URL.html:33
	}
//line views/components/view/URL.html:34
}

//line views/components/view/URL.html:34
func WriteURL(qq422016 qtio422016.Writer, u any, content string, includeExternalIcon bool, ps *cutil.PageState) {
//line views/components/view/URL.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/URL.html:34
	StreamURL(qw422016, u, content, includeExternalIcon, ps)
//line views/components/view/URL.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/URL.html:34
}

//line views/components/view/URL.html:34
func URL(u any, content string, includeExternalIcon bool, ps *cutil.PageState) string {
//line views/components/view/URL.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/URL.html:34
	WriteURL(qb422016, u, content, includeExternalIcon, ps)
//line views/components/view/URL.html:34
	qs422016 := string(qb422016.B)
//line views/components/view/URL.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/URL.html:34
	return qs422016
//line views/components/view/URL.html:34
}

//line views/components/view/URL.html:36
func StreamCodeLink(qw422016 *qt422016.Writer, path string, title string, ps *cutil.PageState) {
//line views/components/view/URL.html:38
	origPath := path
	if title == "" {
		title = path
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	u := util.AppSource + "/blob/master" + path

//line views/components/view/URL.html:47
	StreamURL(qw422016, u, origPath, false, ps)
//line views/components/view/URL.html:48
}

//line views/components/view/URL.html:48
func WriteCodeLink(qq422016 qtio422016.Writer, path string, title string, ps *cutil.PageState) {
//line views/components/view/URL.html:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/URL.html:48
	StreamCodeLink(qw422016, path, title, ps)
//line views/components/view/URL.html:48
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/URL.html:48
}

//line views/components/view/URL.html:48
func CodeLink(path string, title string, ps *cutil.PageState) string {
//line views/components/view/URL.html:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/URL.html:48
	WriteCodeLink(qb422016, path, title, ps)
//line views/components/view/URL.html:48
	qs422016 := string(qb422016.B)
//line views/components/view/URL.html:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/URL.html:48
	return qs422016
//line views/components/view/URL.html:48
}
