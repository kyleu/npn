// Code generated by qtc from "Tree.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vfile/Tree.html:1
package vfile

//line views/vfile/Tree.html:1
import (
	"path"
	"strings"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
)

//line views/vfile/Tree.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/Tree.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/Tree.html:12
func StreamTree(qw422016 *qt422016.Writer, t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Tree.html:13
	components.StreamIndent(qw422016, true, 2)
//line views/vfile/Tree.html:13
	qw422016.N().S(`<ul class="accordion">`)
//line views/vfile/Tree.html:15
	streamtreeNodes(qw422016, t.Nodes, "", urlPrefix, actions, as, ps, 2)
//line views/vfile/Tree.html:16
	components.StreamIndent(qw422016, true, 2)
//line views/vfile/Tree.html:16
	qw422016.N().S(`</ul>`)
//line views/vfile/Tree.html:18
}

//line views/vfile/Tree.html:18
func WriteTree(qq422016 qtio422016.Writer, t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Tree.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:18
	StreamTree(qw422016, t, urlPrefix, actions, as, ps)
//line views/vfile/Tree.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:18
}

//line views/vfile/Tree.html:18
func Tree(t *filesystem.Tree, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/Tree.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:18
	WriteTree(qb422016, t, urlPrefix, actions, as, ps)
//line views/vfile/Tree.html:18
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:18
	return qs422016
//line views/vfile/Tree.html:18
}

//line views/vfile/Tree.html:20
func streamtreeNode(qw422016 *qt422016.Writer, n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:22
	pathID := n.Name
	if pth != "" {
		pathID = pth + "/" + pathID
	}
	pathID = strings.ReplaceAll(pathID, "/", "__")

//line views/vfile/Tree.html:28
	components.StreamIndent(qw422016, true, indent)
//line views/vfile/Tree.html:28
	qw422016.N().S(`<li>`)
//line views/vfile/Tree.html:30
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:30
	qw422016.N().S(`<input id="accordion-`)
//line views/vfile/Tree.html:31
	qw422016.E().S(pathID)
//line views/vfile/Tree.html:31
	qw422016.N().S(`" type="checkbox" hidden="hidden" />`)
//line views/vfile/Tree.html:32
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:32
	qw422016.N().S(`<label for="accordion-`)
//line views/vfile/Tree.html:33
	qw422016.E().S(pathID)
//line views/vfile/Tree.html:33
	qw422016.N().S(`">`)
//line views/vfile/Tree.html:35
	var acts map[string]string
	if actions != nil {
		acts = actions(pth, n)
	}

//line views/vfile/Tree.html:39
	qw422016.N().S(`<div class="right">`)
//line views/vfile/Tree.html:41
	if n.Size > 0 {
//line views/vfile/Tree.html:41
		qw422016.N().S(`<em>`)
//line views/vfile/Tree.html:42
		qw422016.E().S(util.ByteSizeSI(int64(n.Size)))
//line views/vfile/Tree.html:42
		qw422016.N().S(`</em>`)
//line views/vfile/Tree.html:42
		qw422016.N().S(` `)
//line views/vfile/Tree.html:43
	}
//line views/vfile/Tree.html:44
	for k, v := range acts {
//line views/vfile/Tree.html:44
		qw422016.N().S(`<a href="`)
//line views/vfile/Tree.html:45
		qw422016.E().S(v)
//line views/vfile/Tree.html:45
		qw422016.N().S(`">`)
//line views/vfile/Tree.html:45
		qw422016.E().S(k)
//line views/vfile/Tree.html:45
		qw422016.N().S(`</a>`)
//line views/vfile/Tree.html:46
	}
//line views/vfile/Tree.html:46
	qw422016.N().S(`</div>`)
//line views/vfile/Tree.html:48
	components.StreamExpandCollapse(qw422016, indent+2, ps)
//line views/vfile/Tree.html:49
	if n.Dir {
//line views/vfile/Tree.html:50
		components.StreamSVGSimple(qw422016, `folder`, 15, ps)
//line views/vfile/Tree.html:51
	} else {
//line views/vfile/Tree.html:52
		components.StreamSVGSimple(qw422016, `file`, 15, ps)
//line views/vfile/Tree.html:53
	}
//line views/vfile/Tree.html:54
	qw422016.N().S(` `)
//line views/vfile/Tree.html:54
	qw422016.E().S(n.Name)
//line views/vfile/Tree.html:55
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:55
	qw422016.N().S(`</label>`)
//line views/vfile/Tree.html:57
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:57
	qw422016.N().S(`<div class="bd"><div><div>`)
//line views/vfile/Tree.html:59
	if len(n.Children) == 0 {
//line views/vfile/Tree.html:60
		components.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:60
		qw422016.N().S(`<div>`)
//line views/vfile/Tree.html:61
		qw422016.E().S(n.Name)
//line views/vfile/Tree.html:61
		qw422016.N().S(`</div>`)
//line views/vfile/Tree.html:62
	} else {
//line views/vfile/Tree.html:63
		components.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:63
		qw422016.N().S(`<ul class="accordion" style="margin-left:`)
//line views/vfile/Tree.html:64
		qw422016.N().D((indent / 3) * 6)
//line views/vfile/Tree.html:64
		qw422016.N().S(`px; margin-bottom: 0;">`)
//line views/vfile/Tree.html:65
		streamtreeNodes(qw422016, n.Children, path.Join(pth, n.Name), urlPrefix, actions, as, ps, indent+3)
//line views/vfile/Tree.html:66
		components.StreamIndent(qw422016, true, indent+2)
//line views/vfile/Tree.html:66
		qw422016.N().S(`</ul>`)
//line views/vfile/Tree.html:68
	}
//line views/vfile/Tree.html:69
	components.StreamIndent(qw422016, true, indent+1)
//line views/vfile/Tree.html:69
	qw422016.N().S(`</div></div></div>`)
//line views/vfile/Tree.html:71
	components.StreamIndent(qw422016, true, indent)
//line views/vfile/Tree.html:71
	qw422016.N().S(`</li>`)
//line views/vfile/Tree.html:73
}

//line views/vfile/Tree.html:73
func writetreeNode(qq422016 qtio422016.Writer, n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:73
	streamtreeNode(qw422016, n, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:73
}

//line views/vfile/Tree.html:73
func treeNode(n *filesystem.Node, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) string {
//line views/vfile/Tree.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:73
	writetreeNode(qb422016, n, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:73
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:73
	return qs422016
//line views/vfile/Tree.html:73
}

//line views/vfile/Tree.html:75
func streamtreeNodes(qw422016 *qt422016.Writer, nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:76
	for _, n := range nodes {
//line views/vfile/Tree.html:77
		streamtreeNode(qw422016, n, pth, urlPrefix, actions, as, ps, indent+1)
//line views/vfile/Tree.html:78
	}
//line views/vfile/Tree.html:79
}

//line views/vfile/Tree.html:79
func writetreeNodes(qq422016 qtio422016.Writer, nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) {
//line views/vfile/Tree.html:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Tree.html:79
	streamtreeNodes(qw422016, nodes, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:79
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Tree.html:79
}

//line views/vfile/Tree.html:79
func treeNodes(nodes filesystem.Nodes, pth string, urlPrefix string, actions func(p string, n *filesystem.Node) map[string]string, as *app.State, ps *cutil.PageState, indent int) string {
//line views/vfile/Tree.html:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Tree.html:79
	writetreeNodes(qb422016, nodes, pth, urlPrefix, actions, as, ps, indent)
//line views/vfile/Tree.html:79
	qs422016 := string(qb422016.B)
//line views/vfile/Tree.html:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Tree.html:79
	return qs422016
//line views/vfile/Tree.html:79
}
