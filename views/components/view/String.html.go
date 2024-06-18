// Code generated by qtc from "String.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/String.html:1
package view

//line views/components/view/String.html:1
import (
	"strings"

	"github.com/kyleu/npn/app/controller/cutil"
)

//line views/components/view/String.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/String.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/String.html:7
func StreamString(qw422016 *qt422016.Writer, value string, classes ...string) {
//line views/components/view/String.html:8
	if len(classes) == 0 {
//line views/components/view/String.html:9
		qw422016.E().S(value)
//line views/components/view/String.html:10
	} else {
//line views/components/view/String.html:10
		qw422016.N().S(`<span class="`)
//line views/components/view/String.html:11
		qw422016.E().S(strings.Join(classes, ` `))
//line views/components/view/String.html:11
		qw422016.N().S(`">`)
//line views/components/view/String.html:11
		qw422016.E().S(value)
//line views/components/view/String.html:11
		qw422016.N().S(`</span>`)
//line views/components/view/String.html:12
	}
//line views/components/view/String.html:13
}

//line views/components/view/String.html:13
func WriteString(qq422016 qtio422016.Writer, value string, classes ...string) {
//line views/components/view/String.html:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/String.html:13
	StreamString(qw422016, value, classes...)
//line views/components/view/String.html:13
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/String.html:13
}

//line views/components/view/String.html:13
func String(value string, classes ...string) string {
//line views/components/view/String.html:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/String.html:13
	WriteString(qb422016, value, classes...)
//line views/components/view/String.html:13
	qs422016 := string(qb422016.B)
//line views/components/view/String.html:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/String.html:13
	return qs422016
//line views/components/view/String.html:13
}

//line views/components/view/String.html:15
func StreamStringRich(qw422016 *qt422016.Writer, value string, code bool, maxLength int, classes ...string) {
//line views/components/view/String.html:16
	if maxLength > 0 && len(value) > maxLength {
		value = value[:maxLength]
	}

//line views/components/view/String.html:19
	if code {
//line views/components/view/String.html:20
		if len(classes) == 0 {
//line views/components/view/String.html:20
			qw422016.N().S(`<pre>`)
//line views/components/view/String.html:21
			qw422016.E().S(value)
//line views/components/view/String.html:21
			qw422016.N().S(`</pre>`)
//line views/components/view/String.html:22
		} else {
//line views/components/view/String.html:22
			qw422016.N().S(`<pre class="`)
//line views/components/view/String.html:23
			qw422016.E().S(strings.Join(classes, ` `))
//line views/components/view/String.html:23
			qw422016.N().S(`">`)
//line views/components/view/String.html:23
			qw422016.E().S(value)
//line views/components/view/String.html:23
			qw422016.N().S(`</pre>`)
//line views/components/view/String.html:24
		}
//line views/components/view/String.html:25
	} else {
//line views/components/view/String.html:26
		StreamString(qw422016, value, classes...)
//line views/components/view/String.html:27
	}
//line views/components/view/String.html:28
}

//line views/components/view/String.html:28
func WriteStringRich(qq422016 qtio422016.Writer, value string, code bool, maxLength int, classes ...string) {
//line views/components/view/String.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/String.html:28
	StreamStringRich(qw422016, value, code, maxLength, classes...)
//line views/components/view/String.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/String.html:28
}

//line views/components/view/String.html:28
func StringRich(value string, code bool, maxLength int, classes ...string) string {
//line views/components/view/String.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/String.html:28
	WriteStringRich(qb422016, value, code, maxLength, classes...)
//line views/components/view/String.html:28
	qs422016 := string(qb422016.B)
//line views/components/view/String.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/String.html:28
	return qs422016
//line views/components/view/String.html:28
}

//line views/components/view/String.html:30
func StreamStringArray(qw422016 *qt422016.Writer, value []string) {
//line views/components/view/String.html:31
	if len(value) == 0 {
//line views/components/view/String.html:31
		qw422016.N().S(`<em>empty</em>`)
//line views/components/view/String.html:33
	}
//line views/components/view/String.html:35
	maxCount := 5
	display := value
	var extra int
	if len(value) > maxCount {
		extra = len(value) - maxCount
		display = display[:maxCount]
	}

//line views/components/view/String.html:43
	if extra > 0 {
//line views/components/view/String.html:43
		qw422016.N().S(`<span title="`)
//line views/components/view/String.html:43
		qw422016.E().S(strings.Join(value, `, `))
//line views/components/view/String.html:43
		qw422016.N().S(`">`)
//line views/components/view/String.html:43
	}
//line views/components/view/String.html:44
	for idx, v := range display {
//line views/components/view/String.html:45
		if idx > 0 {
//line views/components/view/String.html:45
			qw422016.N().S(`,`)
//line views/components/view/String.html:45
			qw422016.N().S(` `)
//line views/components/view/String.html:45
		}
//line views/components/view/String.html:46
		qw422016.E().S(v)
//line views/components/view/String.html:47
	}
//line views/components/view/String.html:48
	if extra > 0 {
//line views/components/view/String.html:48
		qw422016.N().S(`, <em>and`)
//line views/components/view/String.html:48
		qw422016.N().S(` `)
//line views/components/view/String.html:48
		qw422016.N().D(extra)
//line views/components/view/String.html:48
		qw422016.N().S(` `)
//line views/components/view/String.html:48
		qw422016.N().S(`more...</em>`)
//line views/components/view/String.html:48
	}
//line views/components/view/String.html:49
	if extra > 0 {
//line views/components/view/String.html:49
		qw422016.N().S(`</span>`)
//line views/components/view/String.html:49
	}
//line views/components/view/String.html:50
}

//line views/components/view/String.html:50
func WriteStringArray(qq422016 qtio422016.Writer, value []string) {
//line views/components/view/String.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/String.html:50
	StreamStringArray(qw422016, value)
//line views/components/view/String.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/String.html:50
}

//line views/components/view/String.html:50
func StringArray(value []string) string {
//line views/components/view/String.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/String.html:50
	WriteStringArray(qb422016, value)
//line views/components/view/String.html:50
	qs422016 := string(qb422016.B)
//line views/components/view/String.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/String.html:50
	return qs422016
//line views/components/view/String.html:50
}

//line views/components/view/String.html:53
func StreamFormat(qw422016 *qt422016.Writer, v string, ext string) {
//line views/components/view/String.html:54
	out, err := cutil.FormatLang(v, ext)

//line views/components/view/String.html:55
	if err == nil {
//line views/components/view/String.html:56
		qw422016.N().S(out)
//line views/components/view/String.html:57
	} else {
//line views/components/view/String.html:58
		qw422016.E().S(err.Error())
//line views/components/view/String.html:59
	}
//line views/components/view/String.html:60
}

//line views/components/view/String.html:60
func WriteFormat(qq422016 qtio422016.Writer, v string, ext string) {
//line views/components/view/String.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/String.html:60
	StreamFormat(qw422016, v, ext)
//line views/components/view/String.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/String.html:60
}

//line views/components/view/String.html:60
func Format(v string, ext string) string {
//line views/components/view/String.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/String.html:60
	WriteFormat(qb422016, v, ext)
//line views/components/view/String.html:60
	qs422016 := string(qb422016.B)
//line views/components/view/String.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/String.html:60
	return qs422016
//line views/components/view/String.html:60
}
