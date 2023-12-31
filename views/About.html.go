// Code generated by qtc from "About.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/About.html:2
package views

//line views/About.html:2
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/layout"
)

//line views/About.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/About.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/About.html:10
type About struct{ layout.Basic }

//line views/About.html:12
func (p *About) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:12
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/About.html:14
	components.StreamSVGRefIcon(qw422016, `app`, ps)
//line views/About.html:14
	qw422016.E().S(util.AppName)
//line views/About.html:14
	qw422016.N().S(`</h3>
    <em>v`)
//line views/About.html:15
	qw422016.E().S(as.BuildInfo.Version)
//line views/About.html:15
	qw422016.N().S(`, started `)
//line views/About.html:15
	qw422016.E().S(util.TimeRelative(&as.Started))
//line views/About.html:15
	qw422016.N().S(`</em>
  </div>
  <div class="card">
    <h3>About</h3>
`)
//line views/About.html:19
	qw422016.N().S(`    <p>Coming soon...</p>
`)
//line views/About.html:21
	qw422016.N().S(`  </div>
  `)
//line views/About.html:23
	StreamSourceCode(qw422016)
//line views/About.html:23
	qw422016.N().S(`
  `)
//line views/About.html:24
	StreamFeedback(qw422016)
//line views/About.html:24
	qw422016.N().S(`
`)
//line views/About.html:25
}

//line views/About.html:25
func (p *About) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/About.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:25
	p.StreamBody(qw422016, as, ps)
//line views/About.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:25
}

//line views/About.html:25
func (p *About) Body(as *app.State, ps *cutil.PageState) string {
//line views/About.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:25
	p.WriteBody(qb422016, as, ps)
//line views/About.html:25
	qs422016 := string(qb422016.B)
//line views/About.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:25
	return qs422016
//line views/About.html:25
}

//line views/About.html:27
func StreamSourceCode(qw422016 *qt422016.Writer) {
//line views/About.html:27
	qw422016.N().S(`
<div class="card">
  <h3>Source Code</h3>
  <p>The project is available on <a href="https://github.com/kyleu/npn" target="_blank" rel="noopener noreferrer">GitHub</a></p>
</div>
`)
//line views/About.html:32
}

//line views/About.html:32
func WriteSourceCode(qq422016 qtio422016.Writer) {
//line views/About.html:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:32
	StreamSourceCode(qw422016)
//line views/About.html:32
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:32
}

//line views/About.html:32
func SourceCode() string {
//line views/About.html:32
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:32
	WriteSourceCode(qb422016)
//line views/About.html:32
	qs422016 := string(qb422016.B)
//line views/About.html:32
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:32
	return qs422016
//line views/About.html:32
}

//line views/About.html:34
func StreamFeedback(qw422016 *qt422016.Writer) {
//line views/About.html:34
	qw422016.N().S(`
<div class="card">
  <h3>Feedback</h3>
  <p>For now, email <a href="mailto:npn@kyleu.com">Kyle U</a></p>
</div>
`)
//line views/About.html:39
}

//line views/About.html:39
func WriteFeedback(qq422016 qtio422016.Writer) {
//line views/About.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/About.html:39
	StreamFeedback(qw422016)
//line views/About.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/About.html:39
}

//line views/About.html:39
func Feedback() string {
//line views/About.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/About.html:39
	WriteFeedback(qb422016)
//line views/About.html:39
	qs422016 := string(qb422016.B)
//line views/About.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/About.html:39
	return qs422016
//line views/About.html:39
}
