// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Table.html:2
package components

//line views/components/Table.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/views/vutil"
)

//line views/components/Table.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Table.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Table.html:13
func StreamTableInput(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:14
	id = cutil.CleanID(key, id)

//line views/components/Table.html:14
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:16
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:16
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:17
	qw422016.E().S(id)
//line views/components/Table.html:17
	qw422016.N().S(`"`)
//line views/components/Table.html:17
	streamtitleFor(qw422016, help)
//line views/components/Table.html:17
	qw422016.N().S(`>`)
//line views/components/Table.html:17
	qw422016.E().S(title)
//line views/components/Table.html:17
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:18
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:18
	qw422016.N().S(`<td>`)
//line views/components/Table.html:19
	StreamFormInput(qw422016, key, id, value, help...)
//line views/components/Table.html:19
	qw422016.N().S(`</td>`)
//line views/components/Table.html:20
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:20
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:22
}

//line views/components/Table.html:22
func WriteTableInput(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:22
	StreamTableInput(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:22
}

//line views/components/Table.html:22
func TableInput(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/Table.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:22
	WriteTableInput(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:22
	qs422016 := string(qb422016.B)
//line views/components/Table.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:22
	return qs422016
//line views/components/Table.html:22
}

//line views/components/Table.html:24
func StreamTableInputPassword(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:25
	id = cutil.CleanID(key, id)

//line views/components/Table.html:25
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:27
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:27
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:28
	qw422016.E().S(id)
//line views/components/Table.html:28
	qw422016.N().S(`"`)
//line views/components/Table.html:28
	streamtitleFor(qw422016, help)
//line views/components/Table.html:28
	qw422016.N().S(`>`)
//line views/components/Table.html:28
	qw422016.E().S(title)
//line views/components/Table.html:28
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:29
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:29
	qw422016.N().S(`<td>`)
//line views/components/Table.html:30
	StreamFormInputPassword(qw422016, key, id, value, help...)
//line views/components/Table.html:30
	qw422016.N().S(`</td>`)
//line views/components/Table.html:31
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:31
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:33
}

//line views/components/Table.html:33
func WriteTableInputPassword(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/Table.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:33
	StreamTableInputPassword(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:33
}

//line views/components/Table.html:33
func TableInputPassword(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/Table.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:33
	WriteTableInputPassword(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:33
	qs422016 := string(qb422016.B)
//line views/components/Table.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:33
	return qs422016
//line views/components/Table.html:33
}

//line views/components/Table.html:35
func StreamTableInputNumber(qw422016 *qt422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/Table.html:36
	id = cutil.CleanID(key, id)

//line views/components/Table.html:36
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:38
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:38
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:39
	qw422016.E().S(id)
//line views/components/Table.html:39
	qw422016.N().S(`"`)
//line views/components/Table.html:39
	streamtitleFor(qw422016, help)
//line views/components/Table.html:39
	qw422016.N().S(`>`)
//line views/components/Table.html:39
	qw422016.E().S(title)
//line views/components/Table.html:39
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:40
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:40
	qw422016.N().S(`<td>`)
//line views/components/Table.html:41
	StreamFormInputNumber(qw422016, key, id, value, help...)
//line views/components/Table.html:41
	qw422016.N().S(`</td>`)
//line views/components/Table.html:42
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:42
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:44
}

//line views/components/Table.html:44
func WriteTableInputNumber(qq422016 qtio422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/Table.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:44
	StreamTableInputNumber(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:44
}

//line views/components/Table.html:44
func TableInputNumber(key string, id string, title string, value int, indent int, help ...string) string {
//line views/components/Table.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:44
	WriteTableInputNumber(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:44
	qs422016 := string(qb422016.B)
//line views/components/Table.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:44
	return qs422016
//line views/components/Table.html:44
}

//line views/components/Table.html:46
func StreamTableInputFloat(qw422016 *qt422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/Table.html:47
	id = cutil.CleanID(key, id)

//line views/components/Table.html:47
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:49
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:49
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:50
	qw422016.E().S(id)
//line views/components/Table.html:50
	qw422016.N().S(`"`)
//line views/components/Table.html:50
	streamtitleFor(qw422016, help)
//line views/components/Table.html:50
	qw422016.N().S(`>`)
//line views/components/Table.html:50
	qw422016.E().S(title)
//line views/components/Table.html:50
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:51
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:51
	qw422016.N().S(`<td>`)
//line views/components/Table.html:52
	StreamFormInputFloat(qw422016, key, id, value, help...)
//line views/components/Table.html:52
	qw422016.N().S(`</td>`)
//line views/components/Table.html:53
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:53
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:55
}

//line views/components/Table.html:55
func WriteTableInputFloat(qq422016 qtio422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/Table.html:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:55
	StreamTableInputFloat(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:55
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:55
}

//line views/components/Table.html:55
func TableInputFloat(key string, id string, title string, value float64, indent int, help ...string) string {
//line views/components/Table.html:55
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:55
	WriteTableInputFloat(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:55
	qs422016 := string(qb422016.B)
//line views/components/Table.html:55
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:55
	return qs422016
//line views/components/Table.html:55
}

//line views/components/Table.html:57
func StreamTableInputTimestamp(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:58
	id = cutil.CleanID(key, id)

//line views/components/Table.html:58
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:60
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:60
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:61
	qw422016.E().S(id)
//line views/components/Table.html:61
	qw422016.N().S(`"`)
//line views/components/Table.html:61
	streamtitleFor(qw422016, help)
//line views/components/Table.html:61
	qw422016.N().S(`>`)
//line views/components/Table.html:61
	qw422016.E().S(title)
//line views/components/Table.html:61
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:62
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:62
	qw422016.N().S(`<td>`)
//line views/components/Table.html:63
	StreamFormInputTimestamp(qw422016, key, id, value, help...)
//line views/components/Table.html:63
	qw422016.N().S(`</td>`)
//line views/components/Table.html:64
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:64
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:66
}

//line views/components/Table.html:66
func WriteTableInputTimestamp(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:66
	StreamTableInputTimestamp(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:66
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:66
}

//line views/components/Table.html:66
func TableInputTimestamp(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/Table.html:66
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:66
	WriteTableInputTimestamp(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:66
	qs422016 := string(qb422016.B)
//line views/components/Table.html:66
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:66
	return qs422016
//line views/components/Table.html:66
}

//line views/components/Table.html:68
func StreamTableInputTimestampDay(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:69
	id = cutil.CleanID(key, id)

//line views/components/Table.html:69
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:71
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:71
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:72
	qw422016.E().S(id)
//line views/components/Table.html:72
	qw422016.N().S(`"`)
//line views/components/Table.html:72
	streamtitleFor(qw422016, help)
//line views/components/Table.html:72
	qw422016.N().S(`>`)
//line views/components/Table.html:72
	qw422016.E().S(title)
//line views/components/Table.html:72
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:73
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:73
	qw422016.N().S(`<td>`)
//line views/components/Table.html:74
	StreamFormInputTimestampDay(qw422016, key, id, value, help...)
//line views/components/Table.html:74
	qw422016.N().S(`</td>`)
//line views/components/Table.html:75
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:75
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:77
}

//line views/components/Table.html:77
func WriteTableInputTimestampDay(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/Table.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:77
	StreamTableInputTimestampDay(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:77
}

//line views/components/Table.html:77
func TableInputTimestampDay(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/Table.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:77
	WriteTableInputTimestampDay(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:77
	qs422016 := string(qb422016.B)
//line views/components/Table.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:77
	return qs422016
//line views/components/Table.html:77
}

//line views/components/Table.html:79
func StreamTableInputUUID(qw422016 *qt422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/Table.html:80
	id = cutil.CleanID(key, id)

//line views/components/Table.html:80
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:82
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:82
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:83
	qw422016.E().S(id)
//line views/components/Table.html:83
	qw422016.N().S(`"`)
//line views/components/Table.html:83
	streamtitleFor(qw422016, help)
//line views/components/Table.html:83
	qw422016.N().S(`>`)
//line views/components/Table.html:83
	qw422016.E().S(title)
//line views/components/Table.html:83
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:84
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:84
	qw422016.N().S(`<td>`)
//line views/components/Table.html:85
	StreamFormInputUUID(qw422016, key, id, value, help...)
//line views/components/Table.html:85
	qw422016.N().S(`</td>`)
//line views/components/Table.html:86
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:86
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:88
}

//line views/components/Table.html:88
func WriteTableInputUUID(qq422016 qtio422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/Table.html:88
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:88
	StreamTableInputUUID(qw422016, key, id, title, value, indent, help...)
//line views/components/Table.html:88
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:88
}

//line views/components/Table.html:88
func TableInputUUID(key string, id string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/Table.html:88
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:88
	WriteTableInputUUID(qb422016, key, id, title, value, indent, help...)
//line views/components/Table.html:88
	qs422016 := string(qb422016.B)
//line views/components/Table.html:88
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:88
	return qs422016
//line views/components/Table.html:88
}

//line views/components/Table.html:90
func StreamTableTextarea(qw422016 *qt422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/Table.html:91
	id = cutil.CleanID(key, id)

//line views/components/Table.html:91
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:93
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:93
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:94
	qw422016.E().S(id)
//line views/components/Table.html:94
	qw422016.N().S(`"`)
//line views/components/Table.html:94
	streamtitleFor(qw422016, help)
//line views/components/Table.html:94
	qw422016.N().S(`>`)
//line views/components/Table.html:94
	qw422016.E().S(title)
//line views/components/Table.html:94
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:95
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:95
	qw422016.N().S(`<td>`)
//line views/components/Table.html:96
	StreamFormTextarea(qw422016, key, id, rows, value, help...)
//line views/components/Table.html:96
	qw422016.N().S(`</td>`)
//line views/components/Table.html:97
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:97
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:99
}

//line views/components/Table.html:99
func WriteTableTextarea(qq422016 qtio422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/Table.html:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:99
	StreamTableTextarea(qw422016, key, id, title, rows, value, indent, help...)
//line views/components/Table.html:99
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:99
}

//line views/components/Table.html:99
func TableTextarea(key string, id string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/Table.html:99
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:99
	WriteTableTextarea(qb422016, key, id, title, rows, value, indent, help...)
//line views/components/Table.html:99
	qs422016 := string(qb422016.B)
//line views/components/Table.html:99
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:99
	return qs422016
//line views/components/Table.html:99
}

//line views/components/Table.html:101
func StreamTableSelect(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:102
	id = cutil.CleanID(key, id)

//line views/components/Table.html:102
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:104
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:104
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:105
	qw422016.E().S(id)
//line views/components/Table.html:105
	qw422016.N().S(`"`)
//line views/components/Table.html:105
	streamtitleFor(qw422016, help)
//line views/components/Table.html:105
	qw422016.N().S(`>`)
//line views/components/Table.html:105
	qw422016.E().S(title)
//line views/components/Table.html:105
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:106
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:106
	qw422016.N().S(`<td>`)
//line views/components/Table.html:107
	StreamFormSelect(qw422016, key, id, value, opts, titles, indent)
//line views/components/Table.html:107
	qw422016.N().S(`</td>`)
//line views/components/Table.html:108
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:108
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:110
}

//line views/components/Table.html:110
func WriteTableSelect(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:110
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:110
	StreamTableSelect(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/Table.html:110
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:110
}

//line views/components/Table.html:110
func TableSelect(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:110
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:110
	WriteTableSelect(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/Table.html:110
	qs422016 := string(qb422016.B)
//line views/components/Table.html:110
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:110
	return qs422016
//line views/components/Table.html:110
}

//line views/components/Table.html:112
func StreamTableDatalist(qw422016 *qt422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:113
	id = cutil.CleanID(key, id)

//line views/components/Table.html:113
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:115
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:115
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:116
	qw422016.E().S(id)
//line views/components/Table.html:116
	qw422016.N().S(`"`)
//line views/components/Table.html:116
	streamtitleFor(qw422016, help)
//line views/components/Table.html:116
	qw422016.N().S(`>`)
//line views/components/Table.html:116
	qw422016.E().S(title)
//line views/components/Table.html:116
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:117
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:117
	qw422016.N().S(`<td>`)
//line views/components/Table.html:118
	StreamFormDatalist(qw422016, key, id, value, opts, titles, indent)
//line views/components/Table.html:118
	qw422016.N().S(`</td>`)
//line views/components/Table.html:119
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:119
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:121
}

//line views/components/Table.html:121
func WriteTableDatalist(qq422016 qtio422016.Writer, key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:121
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:121
	StreamTableDatalist(qw422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/Table.html:121
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:121
}

//line views/components/Table.html:121
func TableDatalist(key string, id string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:121
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:121
	WriteTableDatalist(qb422016, key, id, title, value, opts, titles, indent, help...)
//line views/components/Table.html:121
	qs422016 := string(qb422016.B)
//line views/components/Table.html:121
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:121
	return qs422016
//line views/components/Table.html:121
}

//line views/components/Table.html:123
func StreamTableInputTags(qw422016 *qt422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:124
	id = cutil.CleanID(key, id)

//line views/components/Table.html:124
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:126
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:126
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/Table.html:127
	qw422016.E().S(id)
//line views/components/Table.html:127
	qw422016.N().S(`"`)
//line views/components/Table.html:127
	streamtitleFor(qw422016, help)
//line views/components/Table.html:127
	qw422016.N().S(`>`)
//line views/components/Table.html:127
	qw422016.E().S(title)
//line views/components/Table.html:127
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:128
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:128
	qw422016.N().S(`<td>`)
//line views/components/Table.html:129
	StreamFormInputTags(qw422016, key, id, values, ps, help...)
//line views/components/Table.html:129
	qw422016.N().S(`</td>`)
//line views/components/Table.html:130
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:130
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:132
}

//line views/components/Table.html:132
func WriteTableInputTags(qq422016 qtio422016.Writer, key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:132
	StreamTableInputTags(qw422016, key, id, title, values, ps, indent, help...)
//line views/components/Table.html:132
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:132
}

//line views/components/Table.html:132
func TableInputTags(key string, id string, title string, values []string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/Table.html:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:132
	WriteTableInputTags(qb422016, key, id, title, values, ps, indent, help...)
//line views/components/Table.html:132
	qs422016 := string(qb422016.B)
//line views/components/Table.html:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:132
	return qs422016
//line views/components/Table.html:132
}

//line views/components/Table.html:134
func StreamTableRadio(qw422016 *qt422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:134
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:136
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:136
	qw422016.N().S(`<th class="shrink"><label`)
//line views/components/Table.html:137
	streamtitleFor(qw422016, help)
//line views/components/Table.html:137
	qw422016.N().S(`>`)
//line views/components/Table.html:137
	qw422016.E().S(title)
//line views/components/Table.html:137
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:138
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:138
	qw422016.N().S(`<td>`)
//line views/components/Table.html:140
	StreamFormRadio(qw422016, key, value, opts, titles, indent+2)
//line views/components/Table.html:141
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:141
	qw422016.N().S(`</td>`)
//line views/components/Table.html:143
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:143
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:145
}

//line views/components/Table.html:145
func WriteTableRadio(qq422016 qtio422016.Writer, key string, title string, value string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:145
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:145
	StreamTableRadio(qw422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:145
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:145
}

//line views/components/Table.html:145
func TableRadio(key string, title string, value string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:145
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:145
	WriteTableRadio(qb422016, key, title, value, opts, titles, indent, help...)
//line views/components/Table.html:145
	qs422016 := string(qb422016.B)
//line views/components/Table.html:145
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:145
	return qs422016
//line views/components/Table.html:145
}

//line views/components/Table.html:147
func StreamTableBoolean(qw422016 *qt422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/Table.html:148
	StreamTableRadio(qw422016, key, title, fmt.Sprint(value), []string{"true", "false"}, []string{"True", "False"}, indent)
//line views/components/Table.html:149
}

//line views/components/Table.html:149
func WriteTableBoolean(qq422016 qtio422016.Writer, key string, title string, value bool, indent int, help ...string) {
//line views/components/Table.html:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:149
	StreamTableBoolean(qw422016, key, title, value, indent, help...)
//line views/components/Table.html:149
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:149
}

//line views/components/Table.html:149
func TableBoolean(key string, title string, value bool, indent int, help ...string) string {
//line views/components/Table.html:149
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:149
	WriteTableBoolean(qb422016, key, title, value, indent, help...)
//line views/components/Table.html:149
	qs422016 := string(qb422016.B)
//line views/components/Table.html:149
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:149
	return qs422016
//line views/components/Table.html:149
}

//line views/components/Table.html:151
func StreamTableCheckbox(qw422016 *qt422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:151
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:153
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:153
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:154
	qw422016.E().S(title)
//line views/components/Table.html:154
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:155
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:155
	qw422016.N().S(`<td>`)
//line views/components/Table.html:157
	StreamFormCheckbox(qw422016, key, values, opts, titles, indent+2)
//line views/components/Table.html:158
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:158
	qw422016.N().S(`</td>`)
//line views/components/Table.html:160
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:160
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:162
}

//line views/components/Table.html:162
func WriteTableCheckbox(qq422016 qtio422016.Writer, key string, title string, values []string, opts []string, titles []string, indent int, help ...string) {
//line views/components/Table.html:162
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:162
	StreamTableCheckbox(qw422016, key, title, values, opts, titles, indent, help...)
//line views/components/Table.html:162
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:162
}

//line views/components/Table.html:162
func TableCheckbox(key string, title string, values []string, opts []string, titles []string, indent int, help ...string) string {
//line views/components/Table.html:162
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:162
	WriteTableCheckbox(qb422016, key, title, values, opts, titles, indent, help...)
//line views/components/Table.html:162
	qs422016 := string(qb422016.B)
//line views/components/Table.html:162
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:162
	return qs422016
//line views/components/Table.html:162
}

//line views/components/Table.html:164
func StreamTableInputFile(qw422016 *qt422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/Table.html:164
	qw422016.N().S(`<tr><th class="shrink"><label for="`)
//line views/components/Table.html:166
	qw422016.E().S(id)
//line views/components/Table.html:166
	qw422016.N().S(`">`)
//line views/components/Table.html:166
	qw422016.E().S(title)
//line views/components/Table.html:166
	qw422016.N().S(`</label></th><td>`)
//line views/components/Table.html:168
	StreamFormInputFile(qw422016, key, id, label, value)
//line views/components/Table.html:168
	qw422016.N().S(`</td></tr>`)
//line views/components/Table.html:171
}

//line views/components/Table.html:171
func WriteTableInputFile(qq422016 qtio422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/Table.html:171
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:171
	StreamTableInputFile(qw422016, key, id, title, label, value)
//line views/components/Table.html:171
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:171
}

//line views/components/Table.html:171
func TableInputFile(key string, id string, title string, label string, value string) string {
//line views/components/Table.html:171
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:171
	WriteTableInputFile(qb422016, key, id, title, label, value)
//line views/components/Table.html:171
	qs422016 := string(qb422016.B)
//line views/components/Table.html:171
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:171
	return qs422016
//line views/components/Table.html:171
}

//line views/components/Table.html:173
func StreamTableIcons(qw422016 *qt422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:173
	qw422016.N().S(`<tr>`)
//line views/components/Table.html:175
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:175
	qw422016.N().S(`<th class="shrink"><label>`)
//line views/components/Table.html:176
	qw422016.E().S(title)
//line views/components/Table.html:176
	qw422016.N().S(`</label></th>`)
//line views/components/Table.html:177
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/components/Table.html:177
	qw422016.N().S(`<td>`)
//line views/components/Table.html:178
	StreamIconPicker(qw422016, key, value, ps, indent+2)
//line views/components/Table.html:178
	qw422016.N().S(`</td>`)
//line views/components/Table.html:180
	vutil.StreamIndent(qw422016, true, indent)
//line views/components/Table.html:180
	qw422016.N().S(`</tr>`)
//line views/components/Table.html:182
}

//line views/components/Table.html:182
func WriteTableIcons(qq422016 qtio422016.Writer, key string, title string, value string, ps *cutil.PageState, indent int, help ...string) {
//line views/components/Table.html:182
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:182
	StreamTableIcons(qw422016, key, title, value, ps, indent, help...)
//line views/components/Table.html:182
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:182
}

//line views/components/Table.html:182
func TableIcons(key string, title string, value string, ps *cutil.PageState, indent int, help ...string) string {
//line views/components/Table.html:182
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:182
	WriteTableIcons(qb422016, key, title, value, ps, indent, help...)
//line views/components/Table.html:182
	qs422016 := string(qb422016.B)
//line views/components/Table.html:182
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:182
	return qs422016
//line views/components/Table.html:182
}

//line views/components/Table.html:184
func streamtitleFor(qw422016 *qt422016.Writer, help []string) {
//line views/components/Table.html:185
	if len(help) > 0 {
//line views/components/Table.html:185
		qw422016.N().S(` `)
//line views/components/Table.html:185
		qw422016.N().S(`title="`)
//line views/components/Table.html:185
		qw422016.E().S(strings.Join(help, "; "))
//line views/components/Table.html:185
		qw422016.N().S(`"`)
//line views/components/Table.html:185
	}
//line views/components/Table.html:186
}

//line views/components/Table.html:186
func writetitleFor(qq422016 qtio422016.Writer, help []string) {
//line views/components/Table.html:186
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Table.html:186
	streamtitleFor(qw422016, help)
//line views/components/Table.html:186
	qt422016.ReleaseWriter(qw422016)
//line views/components/Table.html:186
}

//line views/components/Table.html:186
func titleFor(help []string) string {
//line views/components/Table.html:186
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Table.html:186
	writetitleFor(qb422016, help)
//line views/components/Table.html:186
	qs422016 := string(qb422016.B)
//line views/components/Table.html:186
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Table.html:186
	return qs422016
//line views/components/Table.html:186
}
