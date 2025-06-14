// Code generated by qtc from "Form.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsearch/Form.html:1
package vsearch

//line views/vsearch/Form.html:1
import (
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
)

//line views/vsearch/Form.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsearch/Form.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsearch/Form.html:7
func StreamForm(qw422016 *qt422016.Writer, act string, q string, placeholder string, currTags []string, ps *cutil.PageState) {
//line views/vsearch/Form.html:7
	qw422016.N().S(`
`)
//line views/vsearch/Form.html:8
	if placeholder == "" {
		placeholder = "Search"
	}

//line views/vsearch/Form.html:10
	qw422016.N().S(`  <form action="`)
//line views/vsearch/Form.html:11
	qw422016.E().S(act)
//line views/vsearch/Form.html:11
	qw422016.N().S(`" method="get">
`)
//line views/vsearch/Form.html:12
	if len(currTags) > 0 {
//line views/vsearch/Form.html:12
		qw422016.N().S(`    <input type="hidden" name="tags" value="`)
//line views/vsearch/Form.html:13
		qw422016.E().S(util.StringJoin(currTags, `,`))
//line views/vsearch/Form.html:13
		qw422016.N().S(`" />
`)
//line views/vsearch/Form.html:14
	}
//line views/vsearch/Form.html:14
	qw422016.N().S(`    <div class="right">
      <button class="right" type="submit">`)
//line views/vsearch/Form.html:16
	components.StreamSVGRef(qw422016, "search", 20, 20, `icon search-icon`, ps)
//line views/vsearch/Form.html:16
	qw422016.N().S(`</button>
      <input class="right br0" type="text" name="q" value="`)
//line views/vsearch/Form.html:17
	qw422016.E().S(q)
//line views/vsearch/Form.html:17
	qw422016.N().S(`" placeholder="`)
//line views/vsearch/Form.html:17
	qw422016.E().S(placeholder)
//line views/vsearch/Form.html:17
	qw422016.N().S(`" />
    </div>
  </form>
`)
//line views/vsearch/Form.html:20
}

//line views/vsearch/Form.html:20
func WriteForm(qq422016 qtio422016.Writer, act string, q string, placeholder string, currTags []string, ps *cutil.PageState) {
//line views/vsearch/Form.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsearch/Form.html:20
	StreamForm(qw422016, act, q, placeholder, currTags, ps)
//line views/vsearch/Form.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vsearch/Form.html:20
}

//line views/vsearch/Form.html:20
func Form(act string, q string, placeholder string, currTags []string, ps *cutil.PageState) string {
//line views/vsearch/Form.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsearch/Form.html:20
	WriteForm(qb422016, act, q, placeholder, currTags, ps)
//line views/vsearch/Form.html:20
	qs422016 := string(qb422016.B)
//line views/vsearch/Form.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsearch/Form.html:20
	return qs422016
//line views/vsearch/Form.html:20
}
