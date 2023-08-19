// Code generated by qtc from "Code.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldview/Code.html:2
package fieldview

//line views/components/fieldview/Code.html:2
import (
	"github.com/kyleu/npn/app/controller/cutil"
)

//line views/components/fieldview/Code.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/Code.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/Code.html:6
func StreamCode(qw422016 *qt422016.Writer, content string, lang string) {
//line views/components/fieldview/Code.html:7
	out, err := cutil.FormatLang(content, lang)

//line views/components/fieldview/Code.html:8
	if err == nil {
//line views/components/fieldview/Code.html:8
		qw422016.N().S(out)
//line views/components/fieldview/Code.html:8
	} else {
//line views/components/fieldview/Code.html:8
		qw422016.E().S(err.Error())
//line views/components/fieldview/Code.html:8
	}
//line views/components/fieldview/Code.html:9
}

//line views/components/fieldview/Code.html:9
func WriteCode(qq422016 qtio422016.Writer, content string, lang string) {
//line views/components/fieldview/Code.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/Code.html:9
	StreamCode(qw422016, content, lang)
//line views/components/fieldview/Code.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/Code.html:9
}

//line views/components/fieldview/Code.html:9
func Code(content string, lang string) string {
//line views/components/fieldview/Code.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/Code.html:9
	WriteCode(qb422016, content, lang)
//line views/components/fieldview/Code.html:9
	qs422016 := string(qb422016.B)
//line views/components/fieldview/Code.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/Code.html:9
	return qs422016
//line views/components/fieldview/Code.html:9
}
