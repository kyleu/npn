// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Request.html:2
package vadmin

//line views/vadmin/Request.html:2
import (
	"slices"

	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/layout"
)

//line views/vadmin/Request.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Request.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Request.html:14
type Request struct {
	layout.Basic
	RC *fasthttp.RequestCtx
}

//line views/vadmin/Request.html:19
func (p *Request) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-ps"><button type="button">Page State</button></a></div>
    <h3>Request Debug</h3>
    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>ID</td>
            <td>`)
//line views/vadmin/Request.html:34
	qw422016.N().DUL(p.RC.ID())
//line views/vadmin/Request.html:34
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>URL</td>
            <td>`)
//line views/vadmin/Request.html:38
	qw422016.E().S(p.RC.URI().String())
//line views/vadmin/Request.html:38
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Protocol</td>
            <td>`)
//line views/vadmin/Request.html:42
	qw422016.E().S(string(p.RC.Request.URI().Scheme()))
//line views/vadmin/Request.html:42
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Host</td>
            <td>`)
//line views/vadmin/Request.html:46
	qw422016.E().S(string(p.RC.Request.URI().Host()))
//line views/vadmin/Request.html:46
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Path</td>
            <td>`)
//line views/vadmin/Request.html:50
	qw422016.E().S(string(p.RC.Request.URI().Path()))
//line views/vadmin/Request.html:50
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Query String</td>
            <td>`)
//line views/vadmin/Request.html:54
	qw422016.E().S(string(p.RC.Request.URI().QueryString()))
//line views/vadmin/Request.html:54
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Body Size</td>
            <td>`)
//line views/vadmin/Request.html:58
	qw422016.N().D(len(p.RC.Request.Body()))
//line views/vadmin/Request.html:58
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>Browser</td>
            <td>`)
//line views/vadmin/Request.html:62
	qw422016.E().S(ps.Browser)
//line views/vadmin/Request.html:62
	qw422016.N().S(` `)
//line views/vadmin/Request.html:62
	qw422016.E().S(ps.BrowserVersion)
//line views/vadmin/Request.html:62
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <td>OS</td>
            <td>`)
//line views/vadmin/Request.html:66
	qw422016.E().S(ps.OS)
//line views/vadmin/Request.html:66
	qw422016.N().S(` `)
//line views/vadmin/Request.html:66
	qw422016.E().S(ps.OSVersion)
//line views/vadmin/Request.html:66
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:72
	if p.RC.Request.Header.Len() > 0 {
//line views/vadmin/Request.html:74
		hd := cutil.RequestHeadersMap(p.RC)

//line views/vadmin/Request.html:75
		qw422016.N().S(`  <div class="card">
    <h3>Headers</h3>
    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vadmin/Request.html:88
		keys := lo.Keys(hd)
		slices.Sort(keys)

//line views/vadmin/Request.html:91
		for _, k := range keys {
//line views/vadmin/Request.html:91
			qw422016.N().S(`          <tr>
            <td class="nowrap">`)
//line views/vadmin/Request.html:93
			qw422016.E().S(k)
//line views/vadmin/Request.html:93
			qw422016.N().S(`</td>
            <td>`)
//line views/vadmin/Request.html:94
			qw422016.E().S(hd[k])
//line views/vadmin/Request.html:94
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vadmin/Request.html:96
		}
//line views/vadmin/Request.html:96
		qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/Request.html:101
	}
//line views/vadmin/Request.html:101
	qw422016.N().S(`  `)
//line views/vadmin/Request.html:102
	components.StreamJSONModal(qw422016, "ps", "Page State", ps, 1)
//line views/vadmin/Request.html:102
	qw422016.N().S(`
`)
//line views/vadmin/Request.html:103
}

//line views/vadmin/Request.html:103
func (p *Request) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Request.html:103
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Request.html:103
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Request.html:103
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Request.html:103
}

//line views/vadmin/Request.html:103
func (p *Request) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Request.html:103
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Request.html:103
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Request.html:103
	qs422016 := string(qb422016.B)
//line views/vadmin/Request.html:103
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Request.html:103
	return qs422016
//line views/vadmin/Request.html:103
}
