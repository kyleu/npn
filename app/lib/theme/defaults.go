// Package theme - Content managed by Project Forge, see [projectforge.md] for details.
package theme

import (
	"fmt"
	"image/color"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/util"
)

const (
	white, black = "#ffffff", "#000000"
	threshold    = (65535 * 3) / 2
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

func TextColorFor(clr string) string {
	c, err := ParseHexColor(clr)
	if err != nil {
		return white
	}
	r, g, b, _ := c.RGBA()
	total := r + g + b
	if total < threshold {
		return white
	}
	return black
}

func ParseHexColor(s string) (color.RGBA, error) {
	ret := color.RGBA{A: 0xff}
	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &ret.R, &ret.G, &ret.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &ret.R, &ret.G, &ret.B)
		// Double the hex digits:
		ret.R *= 17
		ret.G *= 17
		ret.B *= 17
	default:
		err = errors.Errorf("invalid length [%d], must be 7 or 4", len(s))
	}
	return ret, err
}
