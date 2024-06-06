// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/List.html:2
package vtheme

//line views/vtheme/List.html:2
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/theme"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/layout"
)

//line views/vtheme/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/List.html:10
type List struct {
	layout.Basic
	Themes theme.Themes
}

//line views/vtheme/List.html:15
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:15
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vtheme/List.html:17
	components.StreamSVGIcon(qw422016, `eye`, ps)
//line views/vtheme/List.html:17
	qw422016.N().S(`Add Theme</h3>
    <div class="mt">
      <a href="/theme/new" title="add new theme"><button>New Theme</button></a>
      <a href="/theme/palette/crayola" title="add new theme"><button>Choose from Crayola colors</button></a>
      <a href="/theme/palette/css" title="add new theme"><button>Choose from CSS colors</button></a>
      <a href="/theme/palette/wikipedia" title="add new theme"><button>Choose from Wikipedia colors</button></a>
      <hr />
      <form action="/theme/color/edit" method="get">
        <div class="mt">
          <input class="left mrs" type="color" name="color" value="`)
//line views/vtheme/List.html:26
	qw422016.E().S(theme.Default.Light.NavBackground)
//line views/vtheme/List.html:26
	qw422016.N().S(`" />
          <button type="submit">Custom Color Theme</button>
        </div>
      </form>
    </div>
  </div>
  <div class="card">
    <h3>`)
//line views/vtheme/List.html:33
	components.StreamSVGIcon(qw422016, `play`, ps)
//line views/vtheme/List.html:33
	qw422016.N().S(`Current Themes</h3>
    <div class="overflow full-width">
      <div class="theme-container mt">
`)
//line views/vtheme/List.html:36
	for _, t := range p.Themes {
//line views/vtheme/List.html:36
		qw422016.N().S(`        <div class="theme-item">
          <a href="/theme/`)
//line views/vtheme/List.html:38
		qw422016.N().U(t.Key)
//line views/vtheme/List.html:38
		qw422016.N().S(`">
            `)
//line views/vtheme/List.html:39
		StreamMockupTheme(qw422016, t, true, "app", 5, ps)
//line views/vtheme/List.html:39
		qw422016.N().S(`
          </a>
        </div>
`)
//line views/vtheme/List.html:42
	}
//line views/vtheme/List.html:42
	qw422016.N().S(`      </div>
    </div>
  </div>
`)
//line views/vtheme/List.html:46
}

//line views/vtheme/List.html:46
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/List.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/List.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/List.html:46
}

//line views/vtheme/List.html:46
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/List.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/List.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/List.html:46
	qs422016 := string(qb422016.B)
//line views/vtheme/List.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/List.html:46
	return qs422016
//line views/vtheme/List.html:46
}
