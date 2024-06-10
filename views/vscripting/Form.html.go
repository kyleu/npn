// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vscripting/Form.html:2
package vscripting

//line views/vscripting/Form.html:2
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/components/edit"
	"github.com/kyleu/npn/views/layout"
)

//line views/vscripting/Form.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vscripting/Form.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vscripting/Form.html:10
type Form struct {
	layout.Basic
	Path    string
	Content string
}

//line views/vscripting/Form.html:16
func (p *Form) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vscripting/Form.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vscripting/Form.html:18
	if p.Path == "" {
//line views/vscripting/Form.html:18
		qw422016.N().S(`    <h3>`)
//line views/vscripting/Form.html:19
		components.StreamSVGIcon(qw422016, `file`, ps)
//line views/vscripting/Form.html:19
		qw422016.N().S(`New Script</h3>
`)
//line views/vscripting/Form.html:20
	} else {
//line views/vscripting/Form.html:20
		qw422016.N().S(`    <div class="right"><a href="/admin/scripting/`)
//line views/vscripting/Form.html:21
		qw422016.N().U(p.Path)
//line views/vscripting/Form.html:21
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete script [`)
//line views/vscripting/Form.html:21
		qw422016.E().S(p.Path)
//line views/vscripting/Form.html:21
		qw422016.N().S(`]?')"><button>`)
//line views/vscripting/Form.html:21
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vscripting/Form.html:21
		qw422016.N().S(`Delete</button></a></div>
    <h3>`)
//line views/vscripting/Form.html:22
		components.StreamSVGIcon(qw422016, `file`, ps)
//line views/vscripting/Form.html:22
		qw422016.N().S(`Script [`)
//line views/vscripting/Form.html:22
		qw422016.E().S(p.Path)
//line views/vscripting/Form.html:22
		qw422016.N().S(`]</h3>
`)
//line views/vscripting/Form.html:23
	}
//line views/vscripting/Form.html:23
	qw422016.N().S(`    <form action="" class="mt" method="post">
      <div class="overflow full-width">
        <table class="mt expanded">
          <tbody>
            `)
//line views/vscripting/Form.html:28
	if p.Path == "" {
//line views/vscripting/Form.html:28
		edit.StreamStringTable(qw422016, "path", "", "Path", p.Path, 5, "Path to script")
//line views/vscripting/Form.html:28
	}
//line views/vscripting/Form.html:28
	qw422016.N().S(`
            `)
//line views/vscripting/Form.html:29
	edit.StreamTextareaTable(qw422016, "content", "", "Content", 12, p.Content, 5, "Script contents")
//line views/vscripting/Form.html:29
	qw422016.N().S(`
            <tr><td colspan="2"><button type="submit">Save Script</button></td></tr>
          </tbody>
        </table>
      </div>
    </form>
  </div>
`)
//line views/vscripting/Form.html:36
}

//line views/vscripting/Form.html:36
func (p *Form) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vscripting/Form.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vscripting/Form.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vscripting/Form.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vscripting/Form.html:36
}

//line views/vscripting/Form.html:36
func (p *Form) Body(as *app.State, ps *cutil.PageState) string {
//line views/vscripting/Form.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vscripting/Form.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vscripting/Form.html:36
	qs422016 := string(qb422016.B)
//line views/vscripting/Form.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vscripting/Form.html:36
	return qs422016
//line views/vscripting/Form.html:36
}
