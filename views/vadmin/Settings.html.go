// Code generated by qtc from "Settings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Settings.html:1
package vadmin

//line views/vadmin/Settings.html:1
import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/user"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/layout"
	"github.com/kyleu/npn/views/vauth"
)

//line views/vadmin/Settings.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Settings.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Settings.html:11
type Settings struct {
	layout.Basic
	Perms user.Permissions
}

//line views/vadmin/Settings.html:16
func (p *Settings) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:16
	qw422016.N().S(`
  <div class="card">
`)
//line views/vadmin/Settings.html:18
	if util.AppSource != "" {
//line views/vadmin/Settings.html:18
		qw422016.N().S(`    <div class="right"><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:19
		qw422016.E().S(util.AppSource)
//line views/vadmin/Settings.html:19
		qw422016.N().S(`"><button>GitHub</button></a></div>
`)
//line views/vadmin/Settings.html:20
	}
//line views/vadmin/Settings.html:20
	qw422016.N().S(`    <h3 title="github:`)
//line views/vadmin/Settings.html:21
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vadmin/Settings.html:21
	qw422016.N().S(`">`)
//line views/vadmin/Settings.html:21
	components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:21
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:21
	qw422016.E().S(util.AppName)
//line views/vadmin/Settings.html:21
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:21
	qw422016.E().S(as.BuildInfo.String())
//line views/vadmin/Settings.html:21
	qw422016.N().S(`</h3>
`)
//line views/vadmin/Settings.html:22
	if util.AppLegal != "" {
//line views/vadmin/Settings.html:22
		qw422016.N().S(`    <div class="mt">`)
//line views/vadmin/Settings.html:23
		qw422016.N().S(util.AppLegal)
//line views/vadmin/Settings.html:23
		qw422016.N().S(`</div>
`)
//line views/vadmin/Settings.html:24
	}
//line views/vadmin/Settings.html:25
	if util.AppURL != "" {
//line views/vadmin/Settings.html:25
		qw422016.N().S(`    <p><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:26
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:26
		qw422016.N().S(`">`)
//line views/vadmin/Settings.html:26
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:26
		qw422016.N().S(`</a></p>
`)
//line views/vadmin/Settings.html:27
	}
//line views/vadmin/Settings.html:27
	qw422016.N().S(`    <em>This page is for the settings of the application. To change your user preferences, such as themes, <a href="/profile">edit your profile</a>.</em>
  </div>

  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:32
	components.StreamSVGIcon(qw422016, `archive`, ps)
//line views/vadmin/Settings.html:32
	qw422016.N().S(` Admin Functions</h3>
    `)
//line views/vadmin/Settings.html:33
	streamsettingsLink(qw422016, "/admin/server", "archive", "App Information", "All sorts of info about the server and runtime", ps)
//line views/vadmin/Settings.html:33
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:34
	streamsettingsLink(qw422016, "/admin/modules", "archive", "Go Modules", "The Go modules used by "+util.AppName, ps)
//line views/vadmin/Settings.html:34
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:35
	streamsettingsLink(qw422016, "/theme", "archive", "Edit Themes", "Configure the design themes available to end users", ps)
//line views/vadmin/Settings.html:35
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:36
	streamsettingsLink(qw422016, "/admin/scripting", "archive", "Script Files", "View and run the JavaScript files managed by the server", ps)
//line views/vadmin/Settings.html:36
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:37
	streamsettingsLink(qw422016, "/admin/logs", "archive", "Recent Logs", "Displays the 100 most recent app logs", ps)
//line views/vadmin/Settings.html:37
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:41
	components.StreamSVGIcon(qw422016, `bolt`, ps)
//line views/vadmin/Settings.html:41
	qw422016.N().S(` HTTP Methods</h3>
    `)
//line views/vadmin/Settings.html:42
	streamsettingsLink(qw422016, "/admin/sitemap", "bolt", "Sitemap", "Displays the HTTP actions that are available, with documentation", ps)
//line views/vadmin/Settings.html:42
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:43
	streamsettingsLink(qw422016, "/admin/routes", "bolt", "HTTP routes", "Enumerates all registered HTTP routes, by method", ps)
//line views/vadmin/Settings.html:43
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:44
	streamsettingsLink(qw422016, "/admin/session", "bolt", "User Session", "View the user session, including all cookies and settings", ps)
//line views/vadmin/Settings.html:44
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:45
	streamsettingsLink(qw422016, "/admin/request", "bolt", "Debug HTTP Request", "Full debug view of an HTTP request from your browser", ps)
//line views/vadmin/Settings.html:45
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:46
	streamsettingsLink(qw422016, "/admin/sockets", "bolt", "Active WebSockets", "Manage the active WebSockets in this server", ps)
//line views/vadmin/Settings.html:46
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:50
	components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:50
	qw422016.N().S(` App Profiling</h3>
    `)
//line views/vadmin/Settings.html:51
	streamsettingsLink(qw422016, "/admin/memusage", "cog", "Memory Usage", "Detailed memory usage statistics for this application", ps)
//line views/vadmin/Settings.html:51
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:52
	streamsettingsLink(qw422016, "/admin/gc", "cog", "Collect Garbage", "Runs the Go garbage collector", ps)
//line views/vadmin/Settings.html:52
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:53
	streamsettingsLink(qw422016, "/admin/heap", "cog", "Write Memory Dump", "Writes a memory dump to <em>./tmp/mem.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:53
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:54
	streamsettingsLink(qw422016, "/admin/cpu/start", "cog", "Start CPU Profile", "Profiles the CPU using <em>./tmp/cpu.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:54
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:55
	streamsettingsLink(qw422016, "/admin/cpu/stop", "cog", "Stop CPU Profile", "Stops the active CPU profile", ps)
//line views/vadmin/Settings.html:55
	qw422016.N().S(`
    <div class="clear"></div>
  </div>

  `)
//line views/vadmin/Settings.html:59
	vauth.StreamAuthentication(qw422016, as, ps)
//line views/vadmin/Settings.html:59
	qw422016.N().S(`

  `)
//line views/vadmin/Settings.html:61
	vauth.StreamPermissions(qw422016, p.Perms, as)
//line views/vadmin/Settings.html:61
	qw422016.N().S(`
`)
//line views/vadmin/Settings.html:62
}

//line views/vadmin/Settings.html:62
func (p *Settings) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:62
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Settings.html:62
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:62
}

//line views/vadmin/Settings.html:62
func (p *Settings) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:62
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:62
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Settings.html:62
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:62
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:62
	return qs422016
//line views/vadmin/Settings.html:62
}

//line views/vadmin/Settings.html:64
func streamsettingsLink(qw422016 *qt422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:64
	qw422016.N().S(`<hr class="clear" /><div class="mts"><a href="`)
//line views/vadmin/Settings.html:67
	qw422016.E().S(href)
//line views/vadmin/Settings.html:67
	qw422016.N().S(`"><strong>`)
//line views/vadmin/Settings.html:67
	qw422016.E().S(title)
//line views/vadmin/Settings.html:67
	qw422016.N().S(`</strong></a><div><em>`)
//line views/vadmin/Settings.html:68
	qw422016.N().S(description)
//line views/vadmin/Settings.html:68
	qw422016.N().S(`</em></div></div>`)
//line views/vadmin/Settings.html:70
}

//line views/vadmin/Settings.html:70
func writesettingsLink(qq422016 qtio422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:70
	streamsettingsLink(qw422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:70
}

//line views/vadmin/Settings.html:70
func settingsLink(href string, icon string, title string, description string, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:70
	writesettingsLink(qb422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:70
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:70
	return qs422016
//line views/vadmin/Settings.html:70
}
