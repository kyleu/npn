// Code generated by qtc from "Debug.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Debug.html:1
package views

//line views/Debug.html:1
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/layout"
)

//line views/Debug.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Debug.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Debug.html:9
type Debug struct{ layout.Basic }

//line views/Debug.html:11
func (p *Debug) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Debug.html:13
	components.StreamSVGIcon(qw422016, `graph`, ps)
//line views/Debug.html:13
	qw422016.N().S(` `)
//line views/Debug.html:13
	qw422016.E().S(ps.Title)
//line views/Debug.html:13
	qw422016.N().S(`</h3>
`)
//line views/Debug.html:14
	if s, ok := ps.Data.(string); ok {
//line views/Debug.html:14
		qw422016.N().S(`    <div class="overflow full-width"><pre class="mt">`)
//line views/Debug.html:15
		qw422016.E().S(s)
//line views/Debug.html:15
		qw422016.N().S(`</pre></div>
`)
//line views/Debug.html:16
	} else {
//line views/Debug.html:17
		if util.ArrayTest(ps.Data) {
//line views/Debug.html:17
			qw422016.N().S(`    <div class="overflow full-width">
      <table class="mt">
        <tbody>
`)
//line views/Debug.html:21
			a := util.ArrayFromAny[any](ps.Data)

//line views/Debug.html:22
			for idx, x := range a {
//line views/Debug.html:23
				if idx < 32 {
//line views/Debug.html:23
					qw422016.N().S(`          <tr>
            <th class="shrink">`)
//line views/Debug.html:25
					qw422016.N().D(idx + 1)
//line views/Debug.html:25
					qw422016.N().S(`</th>
            <td>`)
//line views/Debug.html:26
					qw422016.N().S(components.JSON(x))
//line views/Debug.html:26
					qw422016.N().S(`</td>
          </tr>
`)
//line views/Debug.html:28
				}
//line views/Debug.html:29
			}
//line views/Debug.html:30
			if len(a) > 32 {
//line views/Debug.html:30
				qw422016.N().S(`          <tr>
            <td class="shrink" colspan="2"><em>...and [`)
//line views/Debug.html:32
				qw422016.N().D(len(a) - 32)
//line views/Debug.html:32
				qw422016.N().S(`] more...</em></td>
          </tr>
`)
//line views/Debug.html:34
			}
//line views/Debug.html:34
			qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/Debug.html:38
		} else {
//line views/Debug.html:38
			qw422016.N().S(`    <div class="mt">`)
//line views/Debug.html:39
			qw422016.N().S(components.JSON(ps.Data))
//line views/Debug.html:39
			qw422016.N().S(`</div>
`)
//line views/Debug.html:40
		}
//line views/Debug.html:41
	}
//line views/Debug.html:41
	qw422016.N().S(`  </div>
`)
//line views/Debug.html:43
}

//line views/Debug.html:43
func (p *Debug) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Debug.html:43
	p.StreamBody(qw422016, as, ps)
//line views/Debug.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/Debug.html:43
}

//line views/Debug.html:43
func (p *Debug) Body(as *app.State, ps *cutil.PageState) string {
//line views/Debug.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Debug.html:43
	p.WriteBody(qb422016, as, ps)
//line views/Debug.html:43
	qs422016 := string(qb422016.B)
//line views/Debug.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Debug.html:43
	return qs422016
//line views/Debug.html:43
}
