// Code generated by qtc from "Int.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldedit/Int.html:2
package fieldedit

//line views/components/fieldedit/Int.html:2
import (
	"github.com/kyleu/npn/views/components"
)

//line views/components/fieldedit/Int.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldedit/Int.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldedit/Int.html:6
func StreamInt(qw422016 *qt422016.Writer, x any, k string) {
//line views/components/fieldedit/Int.html:7
	components.StreamFormInputNumber(qw422016, k, "", x)
//line views/components/fieldedit/Int.html:8
}

//line views/components/fieldedit/Int.html:8
func WriteInt(qq422016 qtio422016.Writer, x any, k string) {
//line views/components/fieldedit/Int.html:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldedit/Int.html:8
	StreamInt(qw422016, x, k)
//line views/components/fieldedit/Int.html:8
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldedit/Int.html:8
}

//line views/components/fieldedit/Int.html:8
func Int(x any, k string) string {
//line views/components/fieldedit/Int.html:8
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldedit/Int.html:8
	WriteInt(qb422016, x, k)
//line views/components/fieldedit/Int.html:8
	qs422016 := string(qb422016.B)
//line views/components/fieldedit/Int.html:8
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldedit/Int.html:8
	return qs422016
//line views/components/fieldedit/Int.html:8
}
