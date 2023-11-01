// Code generated by qtc from "Menu.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Menu.html:2
package layout

//line views/layout/Menu.html:2
import (
	"strings"

	"github.com/kyleu/npn/app/controller/cmenu"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/menu"
	"github.com/kyleu/npn/views/components"
)

//line views/layout/Menu.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Menu.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Menu.html:11
func StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:12
	if len(ps.Menu) > 0 {
//line views/layout/Menu.html:12
		qw422016.N().S(`<div class="menu-container">`)
//line views/layout/Menu.html:14
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:14
		qw422016.N().S(`<div class="menu">`)
//line views/layout/Menu.html:16
		components.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:16
		qw422016.N().S(`<menu class="level-0">`)
//line views/layout/Menu.html:18
		for _, i := range ps.Menu {
//line views/layout/Menu.html:19
			StreamMenuItem(qw422016, i, []string{}, ps.Breadcrumbs, 3, ps)
//line views/layout/Menu.html:20
		}
//line views/layout/Menu.html:21
		components.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:21
		qw422016.N().S(`</menu>`)
//line views/layout/Menu.html:23
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:23
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:25
		components.StreamIndent(qw422016, true, 1)
//line views/layout/Menu.html:25
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:27
	}
//line views/layout/Menu.html:28
}

//line views/layout/Menu.html:28
func WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:28
	StreamMenu(qw422016, ps)
//line views/layout/Menu.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:28
}

//line views/layout/Menu.html:28
func Menu(ps *cutil.PageState) string {
//line views/layout/Menu.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:28
	WriteMenu(qb422016, ps)
//line views/layout/Menu.html:28
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:28
	return qs422016
//line views/layout/Menu.html:28
}

//line views/layout/Menu.html:30
func StreamMenuItem(qw422016 *qt422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:32
	path = append(path, i.Key)
	active, final := breadcrumbs.Active(i, path)

//line views/layout/Menu.html:35
	if i.Key == "" {
//line views/layout/Menu.html:36
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:36
		qw422016.N().S(`<li class="separator"></li>`)
//line views/layout/Menu.html:38
	} else if len(i.Children) > 0 {
//line views/layout/Menu.html:39
		itemID := strings.Join(path, "--")

//line views/layout/Menu.html:40
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:41
		if active {
//line views/layout/Menu.html:41
			qw422016.N().S(`<li class="active" data-menu-key="`)
//line views/layout/Menu.html:41
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:41
			qw422016.N().S(`">`)
//line views/layout/Menu.html:41
		} else {
//line views/layout/Menu.html:41
			qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:41
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:41
			qw422016.N().S(`">`)
//line views/layout/Menu.html:41
		}
//line views/layout/Menu.html:42
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:42
		qw422016.N().S(`<input id="`)
//line views/layout/Menu.html:43
		qw422016.E().S(itemID)
//line views/layout/Menu.html:43
		qw422016.N().S(`-input" type="checkbox"`)
//line views/layout/Menu.html:43
		if active {
//line views/layout/Menu.html:43
			qw422016.N().S(` `)
//line views/layout/Menu.html:43
			qw422016.N().S(`checked="checked"`)
//line views/layout/Menu.html:43
		}
//line views/layout/Menu.html:43
		qw422016.N().S(` `)
//line views/layout/Menu.html:43
		qw422016.N().S(`hidden="hidden" />`)
//line views/layout/Menu.html:44
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:45
		if final {
//line views/layout/Menu.html:45
			qw422016.N().S(`<label class="final" for="`)
//line views/layout/Menu.html:45
			qw422016.E().S(itemID)
//line views/layout/Menu.html:45
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:45
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:45
			qw422016.N().S(`">`)
//line views/layout/Menu.html:45
		} else {
//line views/layout/Menu.html:45
			qw422016.N().S(`<label for="`)
//line views/layout/Menu.html:45
			qw422016.E().S(itemID)
//line views/layout/Menu.html:45
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:45
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:45
			qw422016.N().S(`">`)
//line views/layout/Menu.html:45
		}
//line views/layout/Menu.html:46
		if i.Route != "" {
//line views/layout/Menu.html:47
			components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:47
			qw422016.N().S(`<a class="label-link" href="`)
//line views/layout/Menu.html:48
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:48
			qw422016.N().S(`">`)
//line views/layout/Menu.html:48
			components.StreamSVGRef(qw422016, `link`, 15, 15, ``, ps)
//line views/layout/Menu.html:48
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:49
		}
//line views/layout/Menu.html:50
		components.StreamExpandCollapse(qw422016, indent+3, ps)
//line views/layout/Menu.html:51
		if i.Badge != "" {
//line views/layout/Menu.html:52
			components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:52
			qw422016.N().S(`<div class="badge">`)
//line views/layout/Menu.html:53
			qw422016.E().S(i.Badge)
//line views/layout/Menu.html:53
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:54
		}
//line views/layout/Menu.html:55
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:56
		if i.Icon != "" {
//line views/layout/Menu.html:57
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:58
		}
//line views/layout/Menu.html:59
		if i.Route != "" {
//line views/layout/Menu.html:60
			if i.Warning != "" {
//line views/layout/Menu.html:60
				qw422016.N().S(`<a class="link-confirm" data-message="`)
//line views/layout/Menu.html:61
				qw422016.E().S(i.Warning)
//line views/layout/Menu.html:61
				qw422016.N().S(`" href="`)
//line views/layout/Menu.html:61
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:61
				qw422016.N().S(`">`)
//line views/layout/Menu.html:61
				qw422016.E().S(i.Title)
//line views/layout/Menu.html:61
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:62
			} else {
//line views/layout/Menu.html:62
				qw422016.N().S(`<a href="`)
//line views/layout/Menu.html:63
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:63
				qw422016.N().S(`">`)
//line views/layout/Menu.html:63
				qw422016.E().S(i.Title)
//line views/layout/Menu.html:63
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:64
			}
//line views/layout/Menu.html:65
		} else {
//line views/layout/Menu.html:66
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:67
		}
//line views/layout/Menu.html:68
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:68
		qw422016.N().S(`</label>`)
//line views/layout/Menu.html:70
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:70
		qw422016.N().S(`<div class="menu-content level-`)
//line views/layout/Menu.html:71
		qw422016.N().D(len(path))
//line views/layout/Menu.html:71
		qw422016.N().S(`">`)
//line views/layout/Menu.html:72
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:72
		qw422016.N().S(`<menu>`)
//line views/layout/Menu.html:74
		for _, i := range i.Children {
//line views/layout/Menu.html:75
			StreamMenuItem(qw422016, i, path, breadcrumbs, indent+3, ps)
//line views/layout/Menu.html:76
		}
//line views/layout/Menu.html:77
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:77
		qw422016.N().S(`</menu>`)
//line views/layout/Menu.html:79
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:79
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:81
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:81
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:83
	} else {
//line views/layout/Menu.html:85
		finalClass := "item"
		if active {
			finalClass += " active"
		}
		if final {
			finalClass += " final"
		}
		if i.Warning != "" {
			finalClass += " link-confirm"
		}

//line views/layout/Menu.html:96
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:96
		qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:97
		qw422016.E().S(i.Key)
//line views/layout/Menu.html:97
		qw422016.N().S(`">`)
//line views/layout/Menu.html:98
		if i.Warning != "" {
//line views/layout/Menu.html:98
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:99
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:99
			qw422016.N().S(`" data-message="`)
//line views/layout/Menu.html:99
			qw422016.E().S(i.Warning)
//line views/layout/Menu.html:99
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:99
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:99
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:99
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:99
			qw422016.N().S(`">`)
//line views/layout/Menu.html:100
		} else {
//line views/layout/Menu.html:100
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:101
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:101
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:101
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:101
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:101
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:101
			qw422016.N().S(`">`)
//line views/layout/Menu.html:102
		}
//line views/layout/Menu.html:103
		if i.Badge != "" {
//line views/layout/Menu.html:104
			components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:104
			qw422016.N().S(`<div class="badge">`)
//line views/layout/Menu.html:105
			qw422016.E().S(i.Badge)
//line views/layout/Menu.html:105
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:106
		}
//line views/layout/Menu.html:107
		if i.Icon != "" {
//line views/layout/Menu.html:108
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:109
		}
//line views/layout/Menu.html:110
		qw422016.E().S(i.Title)
//line views/layout/Menu.html:110
		qw422016.N().S(`</a></li>`)
//line views/layout/Menu.html:113
	}
//line views/layout/Menu.html:114
}

//line views/layout/Menu.html:114
func WriteMenuItem(qq422016 qtio422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:114
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:114
	StreamMenuItem(qw422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:114
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:114
}

//line views/layout/Menu.html:114
func MenuItem(i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:114
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:114
	WriteMenuItem(qb422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:114
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:114
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:114
	return qs422016
//line views/layout/Menu.html:114
}
