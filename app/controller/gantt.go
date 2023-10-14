package controller

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/request/queryparam"
	"github.com/kyleu/npn/app/util"
)

type ganttSection struct {
	Key   string `json:"key"`
	Group string `json:"group,omitempty"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func Gantt(rc *fasthttp.RequestCtx) {
	Act("gantt", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rowHeight := 24
		sections, completed, mode := parseGanttRequest(rc)
		pc := func(n int) float64 { return math.Floor((float64(n)/float64(completed))*10000) / 100 }

		ret := make([]string, 0, len(sections)+2)
		ap := func(s string) {
			ret = append(ret, s)
		}

		totalHeight := len(sections) * rowHeight
		svgDecl := `<svg version="1.1" baseProfile="full" height="%v" width="100" preserveAspectRatio="none" viewBox="0 0 100 %v" xmlns="http://www.w3.org/2000/svg">`
		ap(fmt.Sprintf(svgDecl, totalHeight, totalHeight))

		lineMsg := `<line x1="%v" y1="0" x2="%v" y2="%v" stroke="#666" stroke-width="0.1" />`
		for idx := 0; idx < 11; idx++ {
			ap(fmt.Sprintf(lineMsg, idx*10, idx*10, totalHeight))
		}

		bg := `<rect x="0" y="%v" width="100" height="%v" fill="transparent">%v</rect>`
		// bgLine := `<line x1="0" y1="%v" x2="100" y2="%v" stroke="#666" stroke-width="0.1" />`
		rect := `<rect x="%v" y="%v" width="%v" height="%v" style="fill: %v">%v</rect>`
		title := "<title>%v: %v (%v%%)\n%v - %v</title>"

		for idx, section := range sections {
			cy := rowHeight * idx
			per := pc(section.End - section.Start)
			start := pc(section.Start)
			startTitle := util.MicrosToMillis(section.Start)
			width := pc(section.End - section.Start)
			endTitle := util.MicrosToMillis(section.End)
			elapsedTitle := util.MicrosToMillis(section.End - section.Start)
			color := colorForSection(section.Key, mode)

			rectTitle := fmt.Sprintf(title, section.Key, elapsedTitle, per, startTitle, endTitle)
			ap(fmt.Sprintf(bg, cy, rowHeight, rectTitle))
			ap(fmt.Sprintf(rect, start, cy, width, rowHeight, color, rectTitle))
		}

		ap("</svg>")
		return cutil.RespondMIME("", "image/svg+xml", "svg", []byte(strings.Join(ret, "")), rc)
	})
}

func parseGanttRequest(rc *fasthttp.RequestCtx) ([]*ganttSection, int, string) {
	qps := queryparam.QueryParamsFromRaw(string(rc.URI().QueryString()))
	width := -1
	mode := "light"
	ret := make([]*ganttSection, 0)
	get := func(k string) *ganttSection {
		for _, s := range ret {
			if s.Key == k {
				return s
			}
		}
		x := &ganttSection{Key: k}
		ret = append(ret, x)
		return x
	}
	for _, qp := range qps {
		if qp.Key == "w" {
			width, _ = strconv.Atoi(qp.Value)
		}
		if qp.Key == "t" {
			mode = qp.Value
		}
		if strings.Contains(qp.Key, ".") {
			k, t := util.StringSplitLast(qp.Key, '.', true)
			sec := get(k)
			switch t {
			case "g":
				sec.Group = qp.Value
			case "s":
				sec.Start, _ = strconv.Atoi(qp.Value)
			case "e":
				sec.End, _ = strconv.Atoi(qp.Value)
			}
		}
	}
	if width == -1 {
		for _, sec := range ret {
			if sec.End > width {
				width = sec.End
			}
		}
	}
	return ret, width, mode
}

var (
	dark       = "dark"
	lightColor = "#397adb"
	darkColor  = "#101e33"
)

func colorForSection(key string, mode string) string {
	switch key {
	case "dns":
		if mode == dark {
			return "#30444e"
		}
		return "#89b6cc"
	case "connect":
		if mode == dark {
			return "#30444e"
		}
		return "#89b6cc"
	case "tls":
		if mode == dark {
			return "#462206"
		}
		return "#c96112"
	case "reqheaders":
		if mode == dark {
			return "#072918"
		}
		return "#177245"
	case "reqbody":
		if mode == dark {
			return "#072918"
		}
		return "#177245"
	case "rspwait":
		if mode == dark {
			return darkColor
		}
		return lightColor
	case "rspheaders":
		if mode == dark {
			return darkColor
		}
		return lightColor
	case "rspbody":
		if mode == dark {
			return darkColor
		}
		return lightColor
	default:
		if mode == dark {
			return darkColor
		}
		return lightColor
	}
}
