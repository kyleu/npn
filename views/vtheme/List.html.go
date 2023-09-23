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
	"github.com/kyleu/npn/views/layout"
)

//line views/vtheme/List.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/List.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/List.html:9
type List struct {
	layout.Basic
	Themes theme.Themes
}

//line views/vtheme/List.html:14
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:14
	qw422016.N().S(`
  <div class="card">
    <h3>Add Theme</h3>
    <div class="mt">
      <a href="/theme/new" title="add new theme"><button>New Theme</button></a>
      <a href="/theme/palette/crayola" title="add new theme"><button>Choose from Crayola colors</button></a>
      <a href="/theme/palette/css" title="add new theme"><button>Choose from CSS colors</button></a>
      <a href="/theme/palette/wikipedia" title="add new theme"><button>Choose from Wikipedia colors</button></a>
      <hr />
      <form action="/theme/color/edit" method="get">
        <div class="mt">
          <input class="left mrs" type="color" name="color" value="`)
//line views/vtheme/List.html:25
	qw422016.E().S(theme.ThemeDefault.Light.NavBackground)
//line views/vtheme/List.html:25
	qw422016.N().S(`" />
          <button type="submit">Custom Color Theme</button>
        </div>
      </form>
    </div>
  </div>
  <div class="card">
    <h3>Current Themes</h3>
    <div class="overflow full-width">
      <div class="theme-container mt">
`)
//line views/vtheme/List.html:35
	for _, t := range p.Themes {
//line views/vtheme/List.html:35
		qw422016.N().S(`        <div class="theme-item">
          <a href="/theme/`)
//line views/vtheme/List.html:37
		qw422016.N().U(t.Key)
//line views/vtheme/List.html:37
		qw422016.N().S(`">
            `)
//line views/vtheme/List.html:38
		StreamMockupTheme(qw422016, t, true, "app", 5, ps)
//line views/vtheme/List.html:38
		qw422016.N().S(`
          </a>
        </div>
`)
//line views/vtheme/List.html:41
	}
//line views/vtheme/List.html:41
	qw422016.N().S(`      </div>
    </div>
  </div>
`)
//line views/vtheme/List.html:45
}

//line views/vtheme/List.html:45
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/List.html:45
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/List.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/List.html:45
}

//line views/vtheme/List.html:45
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/List.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/List.html:45
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/List.html:45
	qs422016 := string(qb422016.B)
//line views/vtheme/List.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/List.html:45
	return qs422016
//line views/vtheme/List.html:45
}
