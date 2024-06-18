// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vscripting/Detail.html:1
package vscripting

//line views/vscripting/Detail.html:1
import (
	"github.com/samber/lo"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/layout"
)

//line views/vscripting/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vscripting/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vscripting/Detail.html:11
type Detail struct {
	layout.Basic
	Path       string
	Script     string
	LoadResult any
	LoadError  error
	Results    map[string]map[string]any
}

//line views/vscripting/Detail.html:20
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vscripting/Detail.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/scripting/`)
//line views/vscripting/Detail.html:22
	qw422016.N().U(p.Path)
//line views/vscripting/Detail.html:22
	qw422016.N().S(`/edit"><button type="button">Edit</button></a></div>
    <h3>`)
//line views/vscripting/Detail.html:23
	components.StreamSVGIcon(qw422016, `file`, ps)
//line views/vscripting/Detail.html:23
	qw422016.N().S(` `)
//line views/vscripting/Detail.html:23
	qw422016.E().S(p.Path)
//line views/vscripting/Detail.html:23
	qw422016.N().S(`</h3>
    <div class="mt">
`)
//line views/vscripting/Detail.html:25
	out, err := cutil.FormatLang(p.Script, "js")

//line views/vscripting/Detail.html:26
	if err == nil {
//line views/vscripting/Detail.html:26
		qw422016.N().S(`      `)
//line views/vscripting/Detail.html:27
		qw422016.N().S(out)
//line views/vscripting/Detail.html:27
		qw422016.N().S(`
`)
//line views/vscripting/Detail.html:28
	} else {
//line views/vscripting/Detail.html:28
		qw422016.N().S(`      `)
//line views/vscripting/Detail.html:29
		qw422016.E().S(err.Error())
//line views/vscripting/Detail.html:29
		qw422016.N().S(`
`)
//line views/vscripting/Detail.html:30
	}
//line views/vscripting/Detail.html:30
	qw422016.N().S(`    </div>
  </div>
  `)
//line views/vscripting/Detail.html:33
	if p.LoadError != nil {
//line views/vscripting/Detail.html:33
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vscripting/Detail.html:35
		components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vscripting/Detail.html:35
		qw422016.N().S(` Load Error</h3>
    <div class="mt error">`)
//line views/vscripting/Detail.html:36
		qw422016.E().S(p.LoadError.Error())
//line views/vscripting/Detail.html:36
		qw422016.N().S(`</div>
  </div>
`)
//line views/vscripting/Detail.html:38
	}
//line views/vscripting/Detail.html:38
	qw422016.N().S(`  `)
//line views/vscripting/Detail.html:39
	if p.LoadResult != nil {
//line views/vscripting/Detail.html:39
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vscripting/Detail.html:41
		components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vscripting/Detail.html:41
		qw422016.N().S(` Load Result</h3>
    <div class="mt">`)
//line views/vscripting/Detail.html:42
		components.StreamJSON(qw422016, p.LoadResult)
//line views/vscripting/Detail.html:42
		qw422016.N().S(`</div>
  </div>
`)
//line views/vscripting/Detail.html:44
	}
//line views/vscripting/Detail.html:44
	qw422016.N().S(`  `)
//line views/vscripting/Detail.html:45
	for _, f := range util.ArraySorted(lo.Keys(p.Results)) {
//line views/vscripting/Detail.html:46
		res := p.Results[f]

//line views/vscripting/Detail.html:46
		qw422016.N().S(`    `)
//line views/vscripting/Detail.html:47
		if len(res) > 0 {
//line views/vscripting/Detail.html:47
			qw422016.N().S(`    <div class="card">
      <h3>`)
//line views/vscripting/Detail.html:49
			components.StreamSVGIcon(qw422016, `play`, ps)
//line views/vscripting/Detail.html:49
			qw422016.N().S(` [`)
//line views/vscripting/Detail.html:49
			qw422016.E().S(f)
//line views/vscripting/Detail.html:49
			qw422016.N().S(`] Example Results</h3>
      <div class="mt">
        <div class="overflow full-width">
          <table class="expanded min-200">
            <thead>
              <tr>
                <th class="shrink">Example</th>
                <th>Result</th>
              </tr>
            </thead>
            <tbody>
`)
//line views/vscripting/Detail.html:60
			for _, k := range util.ArraySorted(lo.Keys(res)) {
//line views/vscripting/Detail.html:61
				v := res[k]

//line views/vscripting/Detail.html:61
				qw422016.N().S(`              <tr>
                <td><pre>`)
//line views/vscripting/Detail.html:63
				qw422016.E().S(k)
//line views/vscripting/Detail.html:63
				qw422016.N().S(`</pre></td>
                <td><pre>`)
//line views/vscripting/Detail.html:64
				qw422016.E().S(util.ToJSONCompact(v))
//line views/vscripting/Detail.html:64
				qw422016.N().S(`</pre></td>
              </tr>
`)
//line views/vscripting/Detail.html:66
			}
//line views/vscripting/Detail.html:66
			qw422016.N().S(`            </tbody>
          </table>
        </div>
      </div>
    </div>
`)
//line views/vscripting/Detail.html:72
		}
//line views/vscripting/Detail.html:73
	}
//line views/vscripting/Detail.html:74
}

//line views/vscripting/Detail.html:74
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vscripting/Detail.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vscripting/Detail.html:74
	p.StreamBody(qw422016, as, ps)
//line views/vscripting/Detail.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vscripting/Detail.html:74
}

//line views/vscripting/Detail.html:74
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vscripting/Detail.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vscripting/Detail.html:74
	p.WriteBody(qb422016, as, ps)
//line views/vscripting/Detail.html:74
	qs422016 := string(qb422016.B)
//line views/vscripting/Detail.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vscripting/Detail.html:74
	return qs422016
//line views/vscripting/Detail.html:74
}
