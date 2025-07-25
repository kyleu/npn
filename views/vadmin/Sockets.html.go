// Code generated by qtc from "Sockets.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Sockets.html:1
package vadmin

//line views/vadmin/Sockets.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/websocket"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/components"
	"github.com/kyleu/npn/views/components/edit"
	"github.com/kyleu/npn/views/layout"
)

//line views/vadmin/Sockets.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Sockets.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Sockets.html:13
type Sockets struct {
	layout.Basic
	Channels    []string
	Connections []*websocket.Connection
	Taps        []uuid.UUID
}

//line views/vadmin/Sockets.html:20
func (p *Sockets) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/sockets/tap"><button>Tap Traffic</button></a></div>
    <h3>Activity Taps</h3>
    <div class="mt">
`)
//line views/vadmin/Sockets.html:25
	if len(p.Taps) == 0 {
//line views/vadmin/Sockets.html:25
		qw422016.N().S(`      <em>no active taps</em>
`)
//line views/vadmin/Sockets.html:27
	} else {
//line views/vadmin/Sockets.html:27
		qw422016.N().S(`      <ul>
`)
//line views/vadmin/Sockets.html:29
		for _, x := range p.Taps {
//line views/vadmin/Sockets.html:29
			qw422016.N().S(`        <li>`)
//line views/vadmin/Sockets.html:30
			qw422016.E().S(x.String())
//line views/vadmin/Sockets.html:30
			qw422016.N().S(`</li>
`)
//line views/vadmin/Sockets.html:31
		}
//line views/vadmin/Sockets.html:31
		qw422016.N().S(`      </ul>
`)
//line views/vadmin/Sockets.html:33
	}
//line views/vadmin/Sockets.html:33
	qw422016.N().S(`    </div>
  </div>
  <div class="card">
    <h3>Channels</h3>
    <div class="mt">
`)
//line views/vadmin/Sockets.html:39
	if len(p.Channels) == 0 {
//line views/vadmin/Sockets.html:39
		qw422016.N().S(`      <em>no active channels</em>
`)
//line views/vadmin/Sockets.html:41
	} else {
//line views/vadmin/Sockets.html:41
		qw422016.N().S(`      <ul>
`)
//line views/vadmin/Sockets.html:43
		for _, x := range p.Channels {
//line views/vadmin/Sockets.html:43
			qw422016.N().S(`        <li><a href="/admin/sockets/chan/`)
//line views/vadmin/Sockets.html:44
			qw422016.E().S(x)
//line views/vadmin/Sockets.html:44
			qw422016.N().S(`">`)
//line views/vadmin/Sockets.html:44
			qw422016.E().S(x)
//line views/vadmin/Sockets.html:44
			qw422016.N().S(`</a></li>
`)
//line views/vadmin/Sockets.html:45
		}
//line views/vadmin/Sockets.html:45
		qw422016.N().S(`      </ul>
`)
//line views/vadmin/Sockets.html:47
	}
//line views/vadmin/Sockets.html:47
	qw422016.N().S(`    </div>
  </div>
  <div class="card">
    <h3>Active Connections</h3>
    <div class="mt">
`)
//line views/vadmin/Sockets.html:53
	if len(p.Connections) == 0 {
//line views/vadmin/Sockets.html:53
		qw422016.N().S(`      <em>no active connections</em>
`)
//line views/vadmin/Sockets.html:55
	} else {
//line views/vadmin/Sockets.html:55
		qw422016.N().S(`      <div class="overflow full-width">
        <table class="expanded">
          <thead>
            <tr>
              <th class="shrink">ID</th>
              <th>User ID</th>
              <th>Profile Name</th>
              <th>Service</th>
              <th>Channels</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
`)
//line views/vadmin/Sockets.html:69
		for _, x := range p.Connections {
//line views/vadmin/Sockets.html:69
			qw422016.N().S(`            <tr>
              <td class="shrink"><a href="/admin/sockets/conn/`)
//line views/vadmin/Sockets.html:71
			qw422016.E().S(x.ID.String())
//line views/vadmin/Sockets.html:71
			qw422016.N().S(`">`)
//line views/vadmin/Sockets.html:71
			qw422016.E().S(x.ID.String())
//line views/vadmin/Sockets.html:71
			qw422016.N().S(`</a></td>
              <td><a href="/admin/db/user/`)
//line views/vadmin/Sockets.html:72
			qw422016.E().S(x.Profile.ID.String())
//line views/vadmin/Sockets.html:72
			qw422016.N().S(`">`)
//line views/vadmin/Sockets.html:72
			qw422016.E().S(x.Profile.ID.String())
//line views/vadmin/Sockets.html:72
			qw422016.N().S(`</a></td>
              <td>`)
//line views/vadmin/Sockets.html:73
			qw422016.E().S(x.Profile.String())
//line views/vadmin/Sockets.html:73
			qw422016.N().S(`</td>
              <td>`)
//line views/vadmin/Sockets.html:74
			qw422016.E().S(x.Svc)
//line views/vadmin/Sockets.html:74
			qw422016.N().S(`</td>
              <td>`)
//line views/vadmin/Sockets.html:75
			qw422016.E().S(util.StringJoin(x.Channels, ", "))
//line views/vadmin/Sockets.html:75
			qw422016.N().S(`</td>
              <td class="shrink"><a href="#modal-conn-`)
//line views/vadmin/Sockets.html:76
			qw422016.E().S(x.ID.String())
//line views/vadmin/Sockets.html:76
			qw422016.N().S(`"><button type="button">JSON</button></a></td>
            </tr>
`)
//line views/vadmin/Sockets.html:78
		}
//line views/vadmin/Sockets.html:78
		qw422016.N().S(`          </tbody>
        </table>
      </div>
`)
//line views/vadmin/Sockets.html:82
	}
//line views/vadmin/Sockets.html:82
	qw422016.N().S(`    </div>
  </div>
`)
//line views/vadmin/Sockets.html:85
	for _, x := range p.Connections {
//line views/vadmin/Sockets.html:85
		qw422016.N().S(`  `)
//line views/vadmin/Sockets.html:86
		components.StreamJSONModal(qw422016, "conn-"+x.ID.String(), "Connection ["+x.ID.String()+"] JSON", x, 1)
//line views/vadmin/Sockets.html:86
		qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:87
	}
//line views/vadmin/Sockets.html:88
}

//line views/vadmin/Sockets.html:88
func (p *Sockets) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:88
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:88
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sockets.html:88
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:88
}

//line views/vadmin/Sockets.html:88
func (p *Sockets) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:88
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:88
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sockets.html:88
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:88
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:88
	return qs422016
//line views/vadmin/Sockets.html:88
}

//line views/vadmin/Sockets.html:90
type Channel struct {
	layout.Basic
	Channel *websocket.Channel
	Members []*websocket.Connection
}

//line views/vadmin/Sockets.html:96
func (p *Channel) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:96
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-channel"><button type="button">JSON</button></a></div>
    <h3>Channel [`)
//line views/vadmin/Sockets.html:99
	qw422016.E().S(p.Channel.Key)
//line views/vadmin/Sockets.html:99
	qw422016.N().S(`]</h3>
  </div>
  `)
//line views/vadmin/Sockets.html:101
	components.StreamJSONModal(qw422016, "channel", "Channel ["+p.Channel.Key+"] JSON", p.Channel, 1)
//line views/vadmin/Sockets.html:101
	qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:102
	for _, m := range p.Members {
//line views/vadmin/Sockets.html:102
		qw422016.N().S(`  `)
//line views/vadmin/Sockets.html:103
		StreamConnectionCard(qw422016, m, as, ps)
//line views/vadmin/Sockets.html:103
		qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:104
	}
//line views/vadmin/Sockets.html:105
}

//line views/vadmin/Sockets.html:105
func (p *Channel) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:105
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:105
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sockets.html:105
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:105
}

//line views/vadmin/Sockets.html:105
func (p *Channel) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:105
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:105
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sockets.html:105
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:105
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:105
	return qs422016
//line views/vadmin/Sockets.html:105
}

//line views/vadmin/Sockets.html:107
type Connection struct {
	layout.Basic
	Connection *websocket.Connection
}

//line views/vadmin/Sockets.html:112
func (p *Connection) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:112
	qw422016.N().S(`
  `)
//line views/vadmin/Sockets.html:113
	StreamConnectionCard(qw422016, p.Connection, as, ps)
//line views/vadmin/Sockets.html:113
	qw422016.N().S(`
  <div class="card">
    <h3>Send Message</h3>
    <form action="/admin/sockets/conn/`)
//line views/vadmin/Sockets.html:116
	qw422016.E().S(p.Connection.ID.String())
//line views/vadmin/Sockets.html:116
	qw422016.N().S(`/send" method="post">
      <div class="overflow full-width">
        <table class="mt expanded">
          <tbody>
            `)
//line views/vadmin/Sockets.html:120
	edit.StreamStringTable(qw422016, "from", "", "From", ps.Profile.ID.String(), 5, "The user this message is from")
//line views/vadmin/Sockets.html:120
	qw422016.N().S(`
            `)
//line views/vadmin/Sockets.html:121
	edit.StreamStringTable(qw422016, "channel", "", "Channel", "", 5, "The channel this message is for")
//line views/vadmin/Sockets.html:121
	qw422016.N().S(`
            `)
//line views/vadmin/Sockets.html:122
	edit.StreamStringTable(qw422016, "cmd", "", "Command", "", 5, "The command for this message")
//line views/vadmin/Sockets.html:122
	qw422016.N().S(`
            `)
//line views/vadmin/Sockets.html:123
	edit.StreamTextareaTable(qw422016, "param", "", "Parameter", 8, "", 5, "JSON object message payload")
//line views/vadmin/Sockets.html:123
	qw422016.N().S(`
            <tr><td colspan="2"><button type="submit">Send</button></td></tr>
          </tbody>
        </table>
      </div>
    </form>
  </div>
`)
//line views/vadmin/Sockets.html:130
}

//line views/vadmin/Sockets.html:130
func (p *Connection) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:130
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:130
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sockets.html:130
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:130
}

//line views/vadmin/Sockets.html:130
func (p *Connection) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:130
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:130
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sockets.html:130
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:130
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:130
	return qs422016
//line views/vadmin/Sockets.html:130
}

//line views/vadmin/Sockets.html:132
func StreamConnectionCard(qw422016 *qt422016.Writer, c *websocket.Connection, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:132
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="#modal-connection-`)
//line views/vadmin/Sockets.html:134
	qw422016.E().S(c.ID.String())
//line views/vadmin/Sockets.html:134
	qw422016.N().S(`"><button type="button">JSON</button></a></div>
    <h3>`)
//line views/vadmin/Sockets.html:135
	qw422016.E().S(c.ID.String())
//line views/vadmin/Sockets.html:135
	qw422016.N().S(` (`)
//line views/vadmin/Sockets.html:135
	qw422016.E().S(c.Username())
//line views/vadmin/Sockets.html:135
	qw422016.N().S(`)</h3>
    <div class="overflow full-width">
      <table class="mt expanded">
        <tbody>
          <tr>
            <th>Connection ID</th>
            <td>`)
//line views/vadmin/Sockets.html:141
	qw422016.E().S(c.ID.String())
//line views/vadmin/Sockets.html:141
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>User ID</th>
            <td><a href="/admin/db/user/`)
//line views/vadmin/Sockets.html:145
	qw422016.E().S(c.Profile.ID.String())
//line views/vadmin/Sockets.html:145
	qw422016.N().S(`">`)
//line views/vadmin/Sockets.html:145
	qw422016.E().S(c.Profile.ID.String())
//line views/vadmin/Sockets.html:145
	qw422016.N().S(`</a></td>
          </tr>
          <tr>
            <th>Name</th>
            <td>`)
//line views/vadmin/Sockets.html:149
	qw422016.E().S(c.Username())
//line views/vadmin/Sockets.html:149
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Theme</th>
            <td>`)
//line views/vadmin/Sockets.html:153
	qw422016.E().S(c.Profile.Theme)
//line views/vadmin/Sockets.html:153
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Service</th>
            <td>`)
//line views/vadmin/Sockets.html:157
	qw422016.E().S(c.Svc)
//line views/vadmin/Sockets.html:157
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Channels</th>
            <td>`)
//line views/vadmin/Sockets.html:161
	qw422016.E().S(util.StringJoin(c.Channels, ", "))
//line views/vadmin/Sockets.html:161
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  `)
//line views/vadmin/Sockets.html:167
	components.StreamJSONModal(qw422016, "connection-"+c.ID.String(), "Connection ["+c.ID.String()+"] JSON", c, 1)
//line views/vadmin/Sockets.html:167
	qw422016.N().S(`
`)
//line views/vadmin/Sockets.html:168
}

//line views/vadmin/Sockets.html:168
func WriteConnectionCard(qq422016 qtio422016.Writer, c *websocket.Connection, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sockets.html:168
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sockets.html:168
	StreamConnectionCard(qw422016, c, as, ps)
//line views/vadmin/Sockets.html:168
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sockets.html:168
}

//line views/vadmin/Sockets.html:168
func ConnectionCard(c *websocket.Connection, as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sockets.html:168
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sockets.html:168
	WriteConnectionCard(qb422016, c, as, ps)
//line views/vadmin/Sockets.html:168
	qs422016 := string(qb422016.B)
//line views/vadmin/Sockets.html:168
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sockets.html:168
	return qs422016
//line views/vadmin/Sockets.html:168
}
