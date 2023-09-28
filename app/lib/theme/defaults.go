// Package theme - Content managed by Project Forge, see [projectforge.md] for details.
package theme

import (
	"github.com/kyleu/npn/app/util"
)

var Default = func() *Theme {
	nbl := "#193441"
	if o := util.GetEnv("app_nav_color_light"); o != "" {
		nbl = o
	}
	nbd := "#2f4f4f"
	if o := util.GetEnv("app_nav_color_dark"); o != "" {
		nbd = o
	}

	return &Theme{
		Key: "default",
		Light: &Colors{
			Border: "1px solid #dddddd", LinkDecoration: "none",
			Foreground: "#333333", ForegroundMuted: "#777777",
			Background: "#fcfff5", BackgroundMuted: "#e5e9e1",
			LinkForeground: "#304141", LinkVisitedForeground: "#222d2d",
			NavForeground: "#dddddd", NavBackground: nbl,
			MenuForeground: "#cccccc", MenuSelectedForeground: "#dddddd",
			MenuBackground: "#3e606f", MenuSelectedBackground: "#3e606f",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#ff0000",
		},
		Dark: &Colors{
			Border: "1px solid #666666", LinkDecoration: "none",
			Foreground: "#ffffff", ForegroundMuted: "#777777",
			Background: "#121212", BackgroundMuted: "#182322",
			LinkForeground: "#7e9191", LinkVisitedForeground: "#a7b5b4",
			NavForeground: "#ffffff", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuSelectedForeground: "#ffffff",
			MenuBackground: "#203131", MenuSelectedBackground: "#2f4f4f",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#ff0000",
		},
	}
}()
