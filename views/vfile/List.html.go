// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/List.html:2
package vfile

//line views/vfile/List.html:2
import (
	"path/filepath"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/views/components"
)

//line views/vfile/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/List.html:11
func StreamList(qw422016 *qt422016.Writer, path []string, files filesystem.FileInfos, fl filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/List.html:11
	qw422016.N().S(`
  <h3><a href="`)
//line views/vfile/List.html:12
	qw422016.E().S(urlPrefix)
//line views/vfile/List.html:12
	qw422016.N().S(`">.</a>`)
//line views/vfile/List.html:12
	for idx, p := range path {
//line views/vfile/List.html:12
		qw422016.N().S(`/<a href="`)
//line views/vfile/List.html:12
		qw422016.E().S(urlPrefix)
//line views/vfile/List.html:12
		qw422016.N().S(`/`)
//line views/vfile/List.html:12
		qw422016.E().S(filepath.Join(path[:idx+1]...))
//line views/vfile/List.html:12
		qw422016.N().S(`">`)
//line views/vfile/List.html:12
		qw422016.E().S(p)
//line views/vfile/List.html:12
		qw422016.N().S(`</a>`)
//line views/vfile/List.html:12
	}
//line views/vfile/List.html:12
	qw422016.N().S(`</h3>
  <div class="mt">
`)
//line views/vfile/List.html:14
	for _, f := range files {
//line views/vfile/List.html:16
		icon := "file"
		if f.IsDir {
			icon = "folder"
		}
		x := []string{urlPrefix}
		x = append(x, path...)
		x = append(x, f.Name)
		u := filepath.Join(x...)

//line views/vfile/List.html:24
		qw422016.N().S(`    <div><a href="`)
//line views/vfile/List.html:25
		qw422016.E().S(u)
//line views/vfile/List.html:25
		qw422016.N().S(`">`)
//line views/vfile/List.html:25
		components.StreamSVGInline(qw422016, icon, 16, ps)
//line views/vfile/List.html:25
		qw422016.E().S(f.Name)
//line views/vfile/List.html:25
		qw422016.N().S(`</a></div>
`)
//line views/vfile/List.html:26
	}
//line views/vfile/List.html:26
	qw422016.N().S(`  </div>
`)
//line views/vfile/List.html:28
}

//line views/vfile/List.html:28
func WriteList(qq422016 qtio422016.Writer, path []string, files filesystem.FileInfos, fl filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/List.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/List.html:28
	StreamList(qw422016, path, files, fl, urlPrefix, as, ps)
//line views/vfile/List.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/List.html:28
}

//line views/vfile/List.html:28
func List(path []string, files filesystem.FileInfos, fl filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/List.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/List.html:28
	WriteList(qb422016, path, files, fl, urlPrefix, as, ps)
//line views/vfile/List.html:28
	qs422016 := string(qb422016.B)
//line views/vfile/List.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/List.html:28
	return qs422016
//line views/vfile/List.html:28
}
