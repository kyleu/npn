// Code generated by qtc from "Icon.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Icon.html:2
package edit

//line views/components/edit/Icon.html:2
import (
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
)

//line views/components/edit/Icon.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Icon.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Icon.html:8
func StreamIconPicker(qw422016 *qt422016.Writer, key string, selected string, ps *cutil.PageState, indent int) {
//line views/components/edit/Icon.html:9
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Icon.html:9
	qw422016.N().S(`<div class="choice">`)
//line views/components/edit/Icon.html:11
	for _, k := range util.SVGIconKeys {
//line views/components/edit/Icon.html:12
		components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:12
		qw422016.N().S(`<label title="`)
//line views/components/edit/Icon.html:13
		qw422016.E().S(k)
//line views/components/edit/Icon.html:13
		qw422016.N().S(`">`)
//line views/components/edit/Icon.html:14
		if k == selected {
//line views/components/edit/Icon.html:14
			qw422016.N().S(`<input type="radio" name="`)
//line views/components/edit/Icon.html:15
			qw422016.E().S(key)
//line views/components/edit/Icon.html:15
			qw422016.N().S(`" value="`)
//line views/components/edit/Icon.html:15
			qw422016.E().S(k)
//line views/components/edit/Icon.html:15
			qw422016.N().S(`" checked="checked" />`)
//line views/components/edit/Icon.html:16
		} else {
//line views/components/edit/Icon.html:16
			qw422016.N().S(`<input type="radio" name="`)
//line views/components/edit/Icon.html:17
			qw422016.E().S(key)
//line views/components/edit/Icon.html:17
			qw422016.N().S(`" value="`)
//line views/components/edit/Icon.html:17
			qw422016.E().S(k)
//line views/components/edit/Icon.html:17
			qw422016.N().S(`" />`)
//line views/components/edit/Icon.html:18
		}
//line views/components/edit/Icon.html:19
		qw422016.N().S(` `)
//line views/components/edit/Icon.html:20
		components.StreamSVGRef(qw422016, k, 48, 48, "", ps)
//line views/components/edit/Icon.html:20
		qw422016.N().S(`</label>`)
//line views/components/edit/Icon.html:22
	}
//line views/components/edit/Icon.html:23
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:23
	qw422016.N().S(`<div class="clear"></div>`)
//line views/components/edit/Icon.html:25
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Icon.html:25
	qw422016.N().S(`</div>`)
//line views/components/edit/Icon.html:27
}

//line views/components/edit/Icon.html:27
func WriteIconPicker(qq422016 qtio422016.Writer, key string, selected string, ps *cutil.PageState, indent int) {
//line views/components/edit/Icon.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Icon.html:27
	StreamIconPicker(qw422016, key, selected, ps, indent)
//line views/components/edit/Icon.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Icon.html:27
}

//line views/components/edit/Icon.html:27
func IconPicker(key string, selected string, ps *cutil.PageState, indent int) string {
//line views/components/edit/Icon.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Icon.html:27
	WriteIconPicker(qb422016, key, selected, ps, indent)
//line views/components/edit/Icon.html:27
	qs422016 := string(qb422016.B)
//line views/components/edit/Icon.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Icon.html:27
	return qs422016
//line views/components/edit/Icon.html:27
}

//line views/components/edit/Icon.html:29
func StreamIconPickerVertical(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/edit/Icon.html:29
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Icon.html:31
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:31
	qw422016.N().S(`<em class="title">`)
//line views/components/edit/Icon.html:32
	qw422016.E().S(title)
//line views/components/edit/Icon.html:32
	qw422016.N().S(`</em>`)
//line views/components/edit/Icon.html:33
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:33
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Icon.html:34
	StreamIconPicker(qw422016, key, value, ps, indent)
//line views/components/edit/Icon.html:34
	qw422016.N().S(`</div>`)
//line views/components/edit/Icon.html:35
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Icon.html:35
	qw422016.N().S(`</div>`)
//line views/components/edit/Icon.html:37
}

//line views/components/edit/Icon.html:37
func WriteIconPickerVertical(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/edit/Icon.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Icon.html:37
	StreamIconPickerVertical(qw422016, key, title, value, ps, indent)
//line views/components/edit/Icon.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Icon.html:37
}

//line views/components/edit/Icon.html:37
func IconPickerVertical(key string, title string, value string, ps *cutil.PageState, indent int) string {
//line views/components/edit/Icon.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Icon.html:37
	WriteIconPickerVertical(qb422016, key, title, value, ps, indent)
//line views/components/edit/Icon.html:37
	qs422016 := string(qb422016.B)
//line views/components/edit/Icon.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Icon.html:37
	return qs422016
//line views/components/edit/Icon.html:37
}

//line views/components/edit/Icon.html:39
func StreamIcons(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/edit/Icon.html:39
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Icon.html:41
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:41
	qw422016.N().S(`<em class="title">`)
//line views/components/edit/Icon.html:42
	qw422016.E().S(title)
//line views/components/edit/Icon.html:42
	qw422016.N().S(`</em>`)
//line views/components/edit/Icon.html:43
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:43
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Icon.html:44
	StreamIconPicker(qw422016, key, value, ps, indent)
//line views/components/edit/Icon.html:44
	qw422016.N().S(`</div>`)
//line views/components/edit/Icon.html:45
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Icon.html:45
	qw422016.N().S(`</div>`)
//line views/components/edit/Icon.html:47
}

//line views/components/edit/Icon.html:47
func WriteIcons(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int) {
//line views/components/edit/Icon.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Icon.html:47
	StreamIcons(qw422016, key, title, value, ps, indent)
//line views/components/edit/Icon.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Icon.html:47
}

//line views/components/edit/Icon.html:47
func Icons(key string, title string, value string, ps *cutil.PageState, indent int) string {
//line views/components/edit/Icon.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Icon.html:47
	WriteIcons(qb422016, key, title, value, ps, indent)
//line views/components/edit/Icon.html:47
	qs422016 := string(qb422016.B)
//line views/components/edit/Icon.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Icon.html:47
	return qs422016
//line views/components/edit/Icon.html:47
}

//line views/components/edit/Icon.html:49
func StreamIconsTable(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/edit/Icon.html:49
	qw422016.N().S(`<tr>`)
//line views/components/edit/Icon.html:51
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:51
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/edit/Icon.html:52
	qw422016.E().S(title)
//line views/components/edit/Icon.html:52
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Icon.html:53
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Icon.html:53
	qw422016.N().S(`<td>`)
//line views/components/edit/Icon.html:54
	StreamIcons(qw422016, key, title, value, ps, indent+2)
//line views/components/edit/Icon.html:54
	qw422016.N().S(`</td>`)
//line views/components/edit/Icon.html:56
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Icon.html:56
	qw422016.N().S(`</tr>`)
//line views/components/edit/Icon.html:58
}

//line views/components/edit/Icon.html:58
func WriteIconsTable(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/edit/Icon.html:58
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Icon.html:58
	StreamIconsTable(qw422016, key, title, value, ps, indent, help...)
//line views/components/edit/Icon.html:58
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Icon.html:58
}

//line views/components/edit/Icon.html:58
func IconsTable(key string, title string, value string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/edit/Icon.html:58
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Icon.html:58
	WriteIconsTable(qb422016, key, title, value, ps, indent, help...)
//line views/components/edit/Icon.html:58
	qs422016 := string(qb422016.B)
//line views/components/edit/Icon.html:58
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Icon.html:58
	return qs422016
//line views/components/edit/Icon.html:58
}