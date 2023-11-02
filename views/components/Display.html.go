// Code generated by qtc from "Display.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Display.html:2
package components

//line views/components/Display.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/util"
)

//line views/components/Display.html:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Display.html:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Display.html:15
func StreamDisplayTimestamp(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:15
	qw422016.N().S(`<span class="nowrap">`)
//line views/components/Display.html:16
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:16
	qw422016.N().S(`</span>`)
//line views/components/Display.html:17
}

//line views/components/Display.html:17
func WriteDisplayTimestamp(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:17
	StreamDisplayTimestamp(qw422016, value)
//line views/components/Display.html:17
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:17
}

//line views/components/Display.html:17
func DisplayTimestamp(value *time.Time) string {
//line views/components/Display.html:17
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:17
	WriteDisplayTimestamp(qb422016, value)
//line views/components/Display.html:17
	qs422016 := string(qb422016.B)
//line views/components/Display.html:17
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:17
	return qs422016
//line views/components/Display.html:17
}

//line views/components/Display.html:19
func StreamDisplayTimestampRelative(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:19
	qw422016.N().S(`<span class="nowrap reltime" data-time="`)
//line views/components/Display.html:20
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:20
	qw422016.N().S(`">`)
//line views/components/Display.html:20
	qw422016.E().S(util.TimeRelative(value))
//line views/components/Display.html:20
	qw422016.N().S(`</span>`)
//line views/components/Display.html:21
}

//line views/components/Display.html:21
func WriteDisplayTimestampRelative(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:21
	StreamDisplayTimestampRelative(qw422016, value)
//line views/components/Display.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:21
}

//line views/components/Display.html:21
func DisplayTimestampRelative(value *time.Time) string {
//line views/components/Display.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:21
	WriteDisplayTimestampRelative(qb422016, value)
//line views/components/Display.html:21
	qs422016 := string(qb422016.B)
//line views/components/Display.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:21
	return qs422016
//line views/components/Display.html:21
}

//line views/components/Display.html:23
func StreamDisplayTimestampDay(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:24
	qw422016.E().S(util.TimeToYMD(value))
//line views/components/Display.html:25
}

//line views/components/Display.html:25
func WriteDisplayTimestampDay(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:25
	StreamDisplayTimestampDay(qw422016, value)
//line views/components/Display.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:25
}

//line views/components/Display.html:25
func DisplayTimestampDay(value *time.Time) string {
//line views/components/Display.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:25
	WriteDisplayTimestampDay(qb422016, value)
//line views/components/Display.html:25
	qs422016 := string(qb422016.B)
//line views/components/Display.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:25
	return qs422016
//line views/components/Display.html:25
}

//line views/components/Display.html:27
func StreamDisplayUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:28
	if value != nil {
//line views/components/Display.html:29
		qw422016.E().S(value.String())
//line views/components/Display.html:30
	}
//line views/components/Display.html:31
}

//line views/components/Display.html:31
func WriteDisplayUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:31
	StreamDisplayUUID(qw422016, value)
//line views/components/Display.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:31
}

//line views/components/Display.html:31
func DisplayUUID(value *uuid.UUID) string {
//line views/components/Display.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:31
	WriteDisplayUUID(qb422016, value)
//line views/components/Display.html:31
	qs422016 := string(qb422016.B)
//line views/components/Display.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:31
	return qs422016
//line views/components/Display.html:31
}

//line views/components/Display.html:33
func StreamDisplayStringArray(qw422016 *qt422016.Writer, value []string) {
//line views/components/Display.html:34
	if len(value) == 0 {
//line views/components/Display.html:34
		qw422016.N().S(`<em>empty</em>`)
//line views/components/Display.html:36
	}
//line views/components/Display.html:38
	maxCount := 5
	display := value
	var extra int
	if len(value) > maxCount {
		extra = len(value) - maxCount
		display = display[:maxCount]
	}

//line views/components/Display.html:46
	if extra > 0 {
//line views/components/Display.html:46
		qw422016.N().S(`<span title="`)
//line views/components/Display.html:46
		qw422016.E().S(strings.Join(value, `, `))
//line views/components/Display.html:46
		qw422016.N().S(`">`)
//line views/components/Display.html:46
	}
//line views/components/Display.html:47
	for idx, v := range display {
//line views/components/Display.html:48
		if idx > 0 {
//line views/components/Display.html:48
			qw422016.N().S(`,`)
//line views/components/Display.html:48
			qw422016.N().S(` `)
//line views/components/Display.html:48
		}
//line views/components/Display.html:49
		qw422016.E().S(v)
//line views/components/Display.html:50
	}
//line views/components/Display.html:51
	if extra > 0 {
//line views/components/Display.html:51
		qw422016.N().S(`, <em>and`)
//line views/components/Display.html:51
		qw422016.N().S(` `)
//line views/components/Display.html:51
		qw422016.N().D(extra)
//line views/components/Display.html:51
		qw422016.N().S(` `)
//line views/components/Display.html:51
		qw422016.N().S(`more...</em>`)
//line views/components/Display.html:51
	}
//line views/components/Display.html:52
	if extra > 0 {
//line views/components/Display.html:52
		qw422016.N().S(`</span>`)
//line views/components/Display.html:52
	}
//line views/components/Display.html:53
}

//line views/components/Display.html:53
func WriteDisplayStringArray(qq422016 qtio422016.Writer, value []string) {
//line views/components/Display.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:53
	StreamDisplayStringArray(qw422016, value)
//line views/components/Display.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:53
}

//line views/components/Display.html:53
func DisplayStringArray(value []string) string {
//line views/components/Display.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:53
	WriteDisplayStringArray(qb422016, value)
//line views/components/Display.html:53
	qs422016 := string(qb422016.B)
//line views/components/Display.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:53
	return qs422016
//line views/components/Display.html:53
}

//line views/components/Display.html:55
func StreamDisplayIntArray(qw422016 *qt422016.Writer, value []int) {
//line views/components/Display.html:56
	StreamDisplayStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/Display.html:57
}

//line views/components/Display.html:57
func WriteDisplayIntArray(qq422016 qtio422016.Writer, value []int) {
//line views/components/Display.html:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:57
	StreamDisplayIntArray(qw422016, value)
//line views/components/Display.html:57
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:57
}

//line views/components/Display.html:57
func DisplayIntArray(value []int) string {
//line views/components/Display.html:57
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:57
	WriteDisplayIntArray(qb422016, value)
//line views/components/Display.html:57
	qs422016 := string(qb422016.B)
//line views/components/Display.html:57
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:57
	return qs422016
//line views/components/Display.html:57
}

//line views/components/Display.html:59
func StreamDisplayFloatArray(qw422016 *qt422016.Writer, value []float64) {
//line views/components/Display.html:60
	StreamDisplayStringArray(qw422016, util.ArrayToStringArray(value))
//line views/components/Display.html:61
}

//line views/components/Display.html:61
func WriteDisplayFloatArray(qq422016 qtio422016.Writer, value []float64) {
//line views/components/Display.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:61
	StreamDisplayFloatArray(qw422016, value)
//line views/components/Display.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:61
}

//line views/components/Display.html:61
func DisplayFloatArray(value []float64) string {
//line views/components/Display.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:61
	WriteDisplayFloatArray(qb422016, value)
//line views/components/Display.html:61
	qs422016 := string(qb422016.B)
//line views/components/Display.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:61
	return qs422016
//line views/components/Display.html:61
}

//line views/components/Display.html:63
func StreamDisplayDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/Display.html:64
	if len(value) == 0 {
//line views/components/Display.html:64
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/Display.html:66
	} else {
//line views/components/Display.html:66
		qw422016.N().S(`<div class="overflow full-width"><div class="overflow full-width"><table class="expanded"><thead><tr><th>Path</th><th>Old</th><th></th><th>New</th></tr></thead><tbody>`)
//line views/components/Display.html:79
		for _, d := range value {
//line views/components/Display.html:79
			qw422016.N().S(`<tr><td style="width: 30%;"><code>`)
//line views/components/Display.html:81
			qw422016.E().S(d.Path)
//line views/components/Display.html:81
			qw422016.N().S(`</code></td><td style="width: 30%;"><code><em>`)
//line views/components/Display.html:82
			qw422016.E().S(d.Old)
//line views/components/Display.html:82
			qw422016.N().S(`</em></code></td><td style="width: 10%;">→</td><td style="width: 30%;"><code class="success">`)
//line views/components/Display.html:84
			qw422016.E().S(d.New)
//line views/components/Display.html:84
			qw422016.N().S(`</code></td></tr>`)
//line views/components/Display.html:86
		}
//line views/components/Display.html:86
		qw422016.N().S(`</tbody></table></div></div>`)
//line views/components/Display.html:91
	}
//line views/components/Display.html:92
}

//line views/components/Display.html:92
func WriteDisplayDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/Display.html:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:92
	StreamDisplayDiffs(qw422016, value)
//line views/components/Display.html:92
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:92
}

//line views/components/Display.html:92
func DisplayDiffs(value util.Diffs) string {
//line views/components/Display.html:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:92
	WriteDisplayDiffs(qb422016, value)
//line views/components/Display.html:92
	qs422016 := string(qb422016.B)
//line views/components/Display.html:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:92
	return qs422016
//line views/components/Display.html:92
}

//line views/components/Display.html:94
func StreamDisplayDiffsSet(qw422016 *qt422016.Writer, key string, value util.DiffsSet, ps *cutil.PageState) {
//line views/components/Display.html:95
	if len(value) == 0 {
//line views/components/Display.html:95
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/Display.html:97
	} else {
//line views/components/Display.html:97
		qw422016.N().S(`<ul class="accordion">`)
//line views/components/Display.html:99
		for idx, k := range util.ArraySorted[string](lo.Keys(value)) {
//line views/components/Display.html:100
			dk, u := util.StringSplitLast(k, '^', true)

//line views/components/Display.html:101
			if idx < 100 {
//line views/components/Display.html:101
				qw422016.N().S(`<li><input id="accordion-`)
//line views/components/Display.html:103
				qw422016.E().S(k)
//line views/components/Display.html:103
				qw422016.N().S(`-`)
//line views/components/Display.html:103
				qw422016.N().D(idx)
//line views/components/Display.html:103
				qw422016.N().S(`" type="checkbox" hidden="hidden" /><label for="accordion-`)
//line views/components/Display.html:104
				qw422016.E().S(k)
//line views/components/Display.html:104
				qw422016.N().S(`-`)
//line views/components/Display.html:104
				qw422016.N().D(idx)
//line views/components/Display.html:104
				qw422016.N().S(`">`)
//line views/components/Display.html:105
				StreamExpandCollapse(qw422016, 3, ps)
//line views/components/Display.html:106
				if u != "" {
//line views/components/Display.html:106
					qw422016.N().S(`<a href="`)
//line views/components/Display.html:106
					qw422016.E().S(u)
//line views/components/Display.html:106
					qw422016.N().S(`">`)
//line views/components/Display.html:106
					qw422016.E().S(dk)
//line views/components/Display.html:106
					qw422016.N().S(`</a>`)
//line views/components/Display.html:106
				} else {
//line views/components/Display.html:106
					qw422016.E().S(dk)
//line views/components/Display.html:106
				}
//line views/components/Display.html:106
				qw422016.N().S(`</label><div class="bd"><div><div>`)
//line views/components/Display.html:109
				StreamDisplayDiffs(qw422016, value[k])
//line views/components/Display.html:109
				qw422016.N().S(`</div></div></div></li>`)
//line views/components/Display.html:112
			}
//line views/components/Display.html:113
			if idx == 100 {
//line views/components/Display.html:113
				qw422016.N().S(`<li><input id="accordion-`)
//line views/components/Display.html:115
				qw422016.E().S(k)
//line views/components/Display.html:115
				qw422016.N().S(`-extras" type="checkbox" hidden="hidden" /><label for="accordion-`)
//line views/components/Display.html:116
				qw422016.E().S(k)
//line views/components/Display.html:116
				qw422016.N().S(`-extras">...and`)
//line views/components/Display.html:116
				qw422016.N().S(` `)
//line views/components/Display.html:116
				qw422016.N().D(len(value) - 100)
//line views/components/Display.html:116
				qw422016.N().S(` `)
//line views/components/Display.html:116
				qw422016.N().S(`extra</label></li>`)
//line views/components/Display.html:118
			}
//line views/components/Display.html:119
		}
//line views/components/Display.html:119
		qw422016.N().S(`</ul>`)
//line views/components/Display.html:121
	}
//line views/components/Display.html:122
}

//line views/components/Display.html:122
func WriteDisplayDiffsSet(qq422016 qtio422016.Writer, key string, value util.DiffsSet, ps *cutil.PageState) {
//line views/components/Display.html:122
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:122
	StreamDisplayDiffsSet(qw422016, key, value, ps)
//line views/components/Display.html:122
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:122
}

//line views/components/Display.html:122
func DisplayDiffsSet(key string, value util.DiffsSet, ps *cutil.PageState) string {
//line views/components/Display.html:122
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:122
	WriteDisplayDiffsSet(qb422016, key, value, ps)
//line views/components/Display.html:122
	qs422016 := string(qb422016.B)
//line views/components/Display.html:122
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:122
	return qs422016
//line views/components/Display.html:122
}

//line views/components/Display.html:124
func StreamDisplayMaps(qw422016 *qt422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:125
	if len(maps) == 0 {
//line views/components/Display.html:125
		qw422016.N().S(`<em>no results</em>`)
//line views/components/Display.html:127
	} else {
//line views/components/Display.html:127
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/Display.html:132
		for _, k := range maps[0].Keys() {
//line views/components/Display.html:133
			StreamTableHeaderSimple(qw422016, "map", k, k, "", params, nil, ps)
//line views/components/Display.html:134
		}
//line views/components/Display.html:134
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/Display.html:138
		for _, m := range maps {
//line views/components/Display.html:138
			qw422016.N().S(`<tr>`)
//line views/components/Display.html:140
			for _, k := range m.Keys() {
//line views/components/Display.html:142
				res := ""
				switch t := m[k].(type) {
				case string:
					res = t
				case []byte:
					res = string(t)
				default:
					res = fmt.Sprint(m[k])
				}

//line views/components/Display.html:152
				if preserveWhitespace {
//line views/components/Display.html:152
					qw422016.N().S(`<td class="prews">`)
//line views/components/Display.html:153
					qw422016.E().S(res)
//line views/components/Display.html:153
					qw422016.N().S(`</td>`)
//line views/components/Display.html:154
				} else {
//line views/components/Display.html:154
					qw422016.N().S(`<td>`)
//line views/components/Display.html:155
					qw422016.E().S(res)
//line views/components/Display.html:155
					qw422016.N().S(`</td>`)
//line views/components/Display.html:156
				}
//line views/components/Display.html:157
			}
//line views/components/Display.html:157
			qw422016.N().S(`</tr>`)
//line views/components/Display.html:159
		}
//line views/components/Display.html:159
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/Display.html:163
	}
//line views/components/Display.html:164
}

//line views/components/Display.html:164
func WriteDisplayMaps(qq422016 qtio422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:164
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:164
	StreamDisplayMaps(qw422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:164
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:164
}

//line views/components/Display.html:164
func DisplayMaps(maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) string {
//line views/components/Display.html:164
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:164
	WriteDisplayMaps(qb422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:164
	qs422016 := string(qb422016.B)
//line views/components/Display.html:164
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:164
	return qs422016
//line views/components/Display.html:164
}

//line views/components/Display.html:166
func StreamFormat(qw422016 *qt422016.Writer, v string, ext string) {
//line views/components/Display.html:167
	out, err := cutil.FormatLang(v, ext)

//line views/components/Display.html:168
	if err == nil {
//line views/components/Display.html:169
		qw422016.N().S(out)
//line views/components/Display.html:170
	} else {
//line views/components/Display.html:171
		qw422016.E().S(err.Error())
//line views/components/Display.html:172
	}
//line views/components/Display.html:173
}

//line views/components/Display.html:173
func WriteFormat(qq422016 qtio422016.Writer, v string, ext string) {
//line views/components/Display.html:173
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:173
	StreamFormat(qw422016, v, ext)
//line views/components/Display.html:173
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:173
}

//line views/components/Display.html:173
func Format(v string, ext string) string {
//line views/components/Display.html:173
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:173
	WriteFormat(qb422016, v, ext)
//line views/components/Display.html:173
	qs422016 := string(qb422016.B)
//line views/components/Display.html:173
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:173
	return qs422016
//line views/components/Display.html:173
}
