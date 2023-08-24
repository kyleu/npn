// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/npn/tools/wasmserver

import (
	"syscall/js"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/cmd"
)

var (
	version = "0.11.18" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	js.Global().Set("goFetch", js.FuncOf(FetchJS))
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}

func FetchJS(this js.Value, args []js.Value) any {
	return "F00!"
}
