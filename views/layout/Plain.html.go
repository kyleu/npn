// Code generated by qtc from "Plain.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Plain.html:2
package layout

//line views/layout/Plain.html:2
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
)

//line views/layout/Plain.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Plain.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Plain.html:9
type Plain struct{}

var _ Page = (*Plain)(nil)

//line views/layout/Plain.html:12
func (p *Plain) StreamHead(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Plain.html:12
	qw422016.N().S(`
  <meta charset="UTF-8">
  <title>`)
//line views/layout/Plain.html:14
	qw422016.E().S(ps.TitleString())
//line views/layout/Plain.html:14
	qw422016.N().S(`</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover">
  `)
//line views/layout/Plain.html:16
	if ps.Description != "" {
//line views/layout/Plain.html:16
		qw422016.N().S(`<meta property="description" content="`)
//line views/layout/Plain.html:16
		qw422016.E().S(ps.Description)
//line views/layout/Plain.html:16
		qw422016.N().S(`">
  `)
//line views/layout/Plain.html:17
	}
//line views/layout/Plain.html:17
	qw422016.N().S(`<meta property="og:title" content="`)
//line views/layout/Plain.html:17
	qw422016.E().S(ps.TitleString())
//line views/layout/Plain.html:17
	qw422016.N().S(`">
  <meta property="og:image" content="/assets/`)
//line views/layout/Plain.html:18
	qw422016.N().U(util.AppKey)
//line views/layout/Plain.html:18
	qw422016.N().S(`.svg">
  <meta property="og:locale" content="en_US">`)
//line views/layout/Plain.html:19
	qw422016.N().S(ps.HeaderContent)
//line views/layout/Plain.html:19
	qw422016.N().S(`
  <link rel="icon" href="/assets/logo.svg" type="image/svg+xml">
`)
//line views/layout/Plain.html:21
}

//line views/layout/Plain.html:21
func (p *Plain) WriteHead(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Plain.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Plain.html:21
	p.StreamHead(qw422016, as, ps)
//line views/layout/Plain.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Plain.html:21
}

//line views/layout/Plain.html:21
func (p *Plain) Head(as *app.State, ps *cutil.PageState) string {
//line views/layout/Plain.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Plain.html:21
	p.WriteHead(qb422016, as, ps)
//line views/layout/Plain.html:21
	qs422016 := string(qb422016.B)
//line views/layout/Plain.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Plain.html:21
	return qs422016
//line views/layout/Plain.html:21
}

//line views/layout/Plain.html:22
func (p *Plain) StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Plain.html:22
}

//line views/layout/Plain.html:22
func (p *Plain) WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Plain.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Plain.html:22
	p.StreamNav(qw422016, as, ps)
//line views/layout/Plain.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Plain.html:22
}

//line views/layout/Plain.html:22
func (p *Plain) Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Plain.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Plain.html:22
	p.WriteNav(qb422016, as, ps)
//line views/layout/Plain.html:22
	qs422016 := string(qb422016.B)
//line views/layout/Plain.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Plain.html:22
	return qs422016
//line views/layout/Plain.html:22
}

//line views/layout/Plain.html:23
func (p *Plain) StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Plain.html:23
}

//line views/layout/Plain.html:23
func (p *Plain) WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Plain.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Plain.html:23
	p.StreamMenu(qw422016, ps)
//line views/layout/Plain.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Plain.html:23
}

//line views/layout/Plain.html:23
func (p *Plain) Menu(ps *cutil.PageState) string {
//line views/layout/Plain.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Plain.html:23
	p.WriteMenu(qb422016, ps)
//line views/layout/Plain.html:23
	qs422016 := string(qb422016.B)
//line views/layout/Plain.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Plain.html:23
	return qs422016
//line views/layout/Plain.html:23
}

//line views/layout/Plain.html:24
func (p *Plain) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Plain.html:24
	qw422016.N().S(`default body`)
//line views/layout/Plain.html:24
}

//line views/layout/Plain.html:24
func (p *Plain) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Plain.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Plain.html:24
	p.StreamBody(qw422016, as, ps)
//line views/layout/Plain.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Plain.html:24
}

//line views/layout/Plain.html:24
func (p *Plain) Body(as *app.State, ps *cutil.PageState) string {
//line views/layout/Plain.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Plain.html:24
	p.WriteBody(qb422016, as, ps)
//line views/layout/Plain.html:24
	qs422016 := string(qb422016.B)
//line views/layout/Plain.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Plain.html:24
	return qs422016
//line views/layout/Plain.html:24
}
