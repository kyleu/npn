package main // import github.com/kyleu/npn

import (
	"context"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/cmd"
)

var (
	version = "0.2.18" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(context.Background(), &app.BuildInfo{Version: version, Commit: commit, Date: date})
}
