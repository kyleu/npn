// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/npn

import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/cmd"
)

var (
	version = "0.0.4" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
