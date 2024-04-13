// Code generated by qtc from "Color.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Color.html:2
package edit

//line views/components/edit/Color.html:2
import "github.com/kyleu/npn/views/components"

//line views/components/edit/Color.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Color.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Color.html:4
func StreamColor(qw422016 *qt422016.Writer, key string, id string, color string, placeholder ...string) {
//line views/components/edit/Color.html:5
	if id == "" {
//line views/components/edit/Color.html:5
		qw422016.N().S(`<input type="color" name="`)
//line views/components/edit/Color.html:6
		qw422016.E().S(key)
//line views/components/edit/Color.html:6
		qw422016.N().S(`" value="`)
//line views/components/edit/Color.html:6
		qw422016.E().S(color)
//line views/components/edit/Color.html:6
		qw422016.N().S(`"`)
//line views/components/edit/Color.html:6
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Color.html:6
		qw422016.N().S(`/>`)
//line views/components/edit/Color.html:7
	} else {
//line views/components/edit/Color.html:7
		qw422016.N().S(`<input id="`)
//line views/components/edit/Color.html:8
		qw422016.E().S(id)
//line views/components/edit/Color.html:8
		qw422016.N().S(`" type="color" name="`)
//line views/components/edit/Color.html:8
		qw422016.E().S(key)
//line views/components/edit/Color.html:8
		qw422016.N().S(`" value="`)
//line views/components/edit/Color.html:8
		qw422016.E().S(color)
//line views/components/edit/Color.html:8
		qw422016.N().S(`"`)
//line views/components/edit/Color.html:8
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Color.html:8
		qw422016.N().S(`/>`)
//line views/components/edit/Color.html:9
	}
//line views/components/edit/Color.html:10
}

//line views/components/edit/Color.html:10
func WriteColor(qq422016 qtio422016.Writer, key string, id string, color string, placeholder ...string) {
//line views/components/edit/Color.html:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Color.html:10
	StreamColor(qw422016, key, id, color, placeholder...)
//line views/components/edit/Color.html:10
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Color.html:10
}

//line views/components/edit/Color.html:10
func Color(key string, id string, color string, placeholder ...string) string {
//line views/components/edit/Color.html:10
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Color.html:10
	WriteColor(qb422016, key, id, color, placeholder...)
//line views/components/edit/Color.html:10
	qs422016 := string(qb422016.B)
//line views/components/edit/Color.html:10
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Color.html:10
	return qs422016
//line views/components/edit/Color.html:10
}

//line views/components/edit/Color.html:12
func StreamColorTable(qw422016 *qt422016.Writer, key string, id string, title string, color string, indent int, help ...string) {
//line views/components/edit/Color.html:12
	qw422016.N().S(`<tr>`)
//line views/components/edit/Color.html:14
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Color.html:14
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Color.html:15
	qw422016.E().S(id)
//line views/components/edit/Color.html:15
	qw422016.N().S(`">`)
//line views/components/edit/Color.html:15
	qw422016.E().S(title)
//line views/components/edit/Color.html:15
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Color.html:16
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Color.html:16
	qw422016.N().S(`<td>`)
//line views/components/edit/Color.html:18
	components.StreamIndent(qw422016, true, indent+2)
//line views/components/edit/Color.html:19
	StreamColor(qw422016, key, id, color, help...)
//line views/components/edit/Color.html:20
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Color.html:20
	qw422016.N().S(`</td>`)
//line views/components/edit/Color.html:22
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Color.html:22
	qw422016.N().S(`</tr>`)
//line views/components/edit/Color.html:24
}

//line views/components/edit/Color.html:24
func WriteColorTable(qq422016 qtio422016.Writer, key string, id string, title string, color string, indent int, help ...string) {
//line views/components/edit/Color.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Color.html:24
	StreamColorTable(qw422016, key, id, title, color, indent, help...)
//line views/components/edit/Color.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Color.html:24
}

//line views/components/edit/Color.html:24
func ColorTable(key string, id string, title string, color string, indent int, help ...string) string {
//line views/components/edit/Color.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Color.html:24
	WriteColorTable(qb422016, key, id, title, color, indent, help...)
//line views/components/edit/Color.html:24
	qs422016 := string(qb422016.B)
//line views/components/edit/Color.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Color.html:24
	return qs422016
//line views/components/edit/Color.html:24
}
