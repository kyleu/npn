// Code generated by qtc from "Args.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vpage/Args.html:1
package vpage

//line views/vpage/Args.html:1
import (
	"fmt"
	"strings"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components/edit"
	"github.com/kyleu/npn/views/layout"
)

//line views/vpage/Args.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vpage/Args.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vpage/Args.html:12
type Args struct {
	layout.Basic
	URL        string
	Directions string
	Results    *util.FieldDescResults
	Hidden     map[string]string
	Warning    string
}

//line views/vpage/Args.html:21
func (p *Args) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:21
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vpage/Args.html:23
	if p.Directions == "" {
//line views/vpage/Args.html:23
		qw422016.N().S(`Enter Data`)
//line views/vpage/Args.html:23
	} else {
//line views/vpage/Args.html:23
		qw422016.E().S(p.Directions)
//line views/vpage/Args.html:23
	}
//line views/vpage/Args.html:23
	qw422016.N().S(`</h3>
`)
//line views/vpage/Args.html:25
	var onsubmit string
	if p.Warning != "" {
		onsubmit = fmt.Sprintf(` onsubmit="return confirm('%s')"`, strings.ReplaceAll(strings.ReplaceAll(p.Warning, "'", "\\'"), "\"", ""))
	}

//line views/vpage/Args.html:29
	qw422016.N().S(`    <form action="`)
//line views/vpage/Args.html:30
	qw422016.E().S(p.URL)
//line views/vpage/Args.html:30
	qw422016.N().S(`" method="get"`)
//line views/vpage/Args.html:30
	qw422016.N().S(onsubmit)
//line views/vpage/Args.html:30
	qw422016.N().S(`>
`)
//line views/vpage/Args.html:31
	for k, v := range p.Hidden {
//line views/vpage/Args.html:31
		qw422016.N().S(`      <input type="hidden" name="`)
//line views/vpage/Args.html:32
		qw422016.E().S(k)
//line views/vpage/Args.html:32
		qw422016.N().S(`" value="`)
//line views/vpage/Args.html:32
		qw422016.E().S(v)
//line views/vpage/Args.html:32
		qw422016.N().S(`" />
`)
//line views/vpage/Args.html:33
	}
//line views/vpage/Args.html:33
	qw422016.N().S(`      `)
//line views/vpage/Args.html:34
	edit.StreamTableEditorNoForm(qw422016, "args", p.Results.FieldDescs, p.Results.Values, "", "", "Submit")
//line views/vpage/Args.html:34
	qw422016.N().S(`
    </form>
  </div>
`)
//line views/vpage/Args.html:37
}

//line views/vpage/Args.html:37
func (p *Args) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vpage/Args.html:37
	p.StreamBody(qw422016, as, ps)
//line views/vpage/Args.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vpage/Args.html:37
}

//line views/vpage/Args.html:37
func (p *Args) Body(as *app.State, ps *cutil.PageState) string {
//line views/vpage/Args.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vpage/Args.html:37
	p.WriteBody(qb422016, as, ps)
//line views/vpage/Args.html:37
	qs422016 := string(qb422016.B)
//line views/vpage/Args.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vpage/Args.html:37
	return qs422016
//line views/vpage/Args.html:37
}
