// Code generated by qtc from "ExpandCollapse.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/ExpandCollapse.html:1
package components

//line views/components/ExpandCollapse.html:1
import "github.com/kyleu/npn/app/controller/cutil"

//line views/components/ExpandCollapse.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/ExpandCollapse.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/ExpandCollapse.html:3
func StreamExpandCollapse(qw422016 *qt422016.Writer, indent int, ps *cutil.PageState) {
//line views/components/ExpandCollapse.html:4
	StreamIndent(qw422016, true, indent)
//line views/components/ExpandCollapse.html:5
	StreamSVGRef(qw422016, `right`, 15, 15, `expand-collapse`, ps)
//line views/components/ExpandCollapse.html:6
}

//line views/components/ExpandCollapse.html:6
func WriteExpandCollapse(qq422016 qtio422016.Writer, indent int, ps *cutil.PageState) {
//line views/components/ExpandCollapse.html:6
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/ExpandCollapse.html:6
	StreamExpandCollapse(qw422016, indent, ps)
//line views/components/ExpandCollapse.html:6
	qt422016.ReleaseWriter(qw422016)
//line views/components/ExpandCollapse.html:6
}

//line views/components/ExpandCollapse.html:6
func ExpandCollapse(indent int, ps *cutil.PageState) string {
//line views/components/ExpandCollapse.html:6
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/ExpandCollapse.html:6
	WriteExpandCollapse(qb422016, indent, ps)
//line views/components/ExpandCollapse.html:6
	qs422016 := string(qb422016.B)
//line views/components/ExpandCollapse.html:6
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/ExpandCollapse.html:6
	return qs422016
//line views/components/ExpandCollapse.html:6
}
