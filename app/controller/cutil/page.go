package cutil

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cmenu"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/lib/menu"
	"github.com/kyleu/npn/app/lib/telemetry"
	"github.com/kyleu/npn/app/lib/theme"
	"github.com/kyleu/npn/app/lib/user"
	dbuser "github.com/kyleu/npn/app/user"
	"github.com/kyleu/npn/app/util"
)

const (
	DefaultSearchPath  = "/search"
	DefaultProfilePath = "/profile"
	defaultIcon        = "app"
)

var (
	defaultRootTitleAppend = util.GetEnv("app_display_name_append")
	defaultRootTitle       = func() string {
		if tmp := util.GetEnv("app_display_name"); tmp != "" {
			return tmp
		}
		return util.AppName
	}()
)

type PageState struct {
	Action         string            `json:"action,omitempty"`
	Title          string            `json:"title,omitempty"`
	Description    string            `json:"description,omitempty"`
	Method         string            `json:"method,omitempty"`
	URI            *url.URL          `json:"-"`
	Menu           menu.Items        `json:"menu,omitempty"`
	Breadcrumbs    cmenu.Breadcrumbs `json:"breadcrumbs,omitempty"`
	Flashes        []string          `json:"flashes,omitempty"`
	Session        util.ValueMap     `json:"-"`
	User           *dbuser.User      `json:"user,omitempty"`
	Profile        *user.Profile     `json:"profile,omitempty"`
	Accounts       user.Accounts     `json:"accounts,omitempty"`
	Authed         bool              `json:"authed,omitempty"`
	Admin          bool              `json:"admin,omitempty"`
	Params         filter.ParamSet   `json:"params,omitempty"`
	Icons          []string          `json:"icons,omitempty"`
	DefaultNavIcon string            `json:"defaultNavIcon,omitempty"`
	RootIcon       string            `json:"rootIcon,omitempty"`
	RootPath       string            `json:"rootPath,omitempty"`
	RootTitle      string            `json:"rootTitle,omitempty"`
	SearchPath     string            `json:"searchPath,omitempty"`
	ProfilePath    string            `json:"profilePath,omitempty"`
	HideHeader     bool              `json:"hideHeader,omitempty"`
	HideMenu       bool              `json:"hideMenu,omitempty"`
	NoScript       bool              `json:"noScript,omitempty"`
	ForceRedirect  string            `json:"forceRedirect,omitempty"`
	DefaultFormat  string            `json:"defaultFormat,omitempty"`
	HeaderContent  string            `json:"headerContent,omitempty"`
	Browser        string            `json:"browser,omitempty"`
	BrowserVersion string            `json:"browserVersion,omitempty"`
	OS             string            `json:"os,omitempty"`
	OSVersion      string            `json:"osVersion,omitempty"`
	Platform       string            `json:"platform,omitempty"`
	Data           any               `json:"data,omitempty"`
	Started        time.Time         `json:"started,omitempty"`
	RenderElapsed  float64           `json:"renderElapsed,omitempty"`
	ResponseBytes  int64             `json:"responseBytes,omitempty"`
	RequestBody    []byte            `json:"-"`
	W              *WriteCounter     `json:"-"`
	Logger         util.Logger       `json:"-"`
	Context        context.Context   `json:"-"` //nolint:containedctx // properly closed, never directly used
	Span           *telemetry.Span   `json:"-"`
}

func (p *PageState) AddIcon(keys ...string) {
	lo.ForEach(keys, func(k string, _ int) {
		if !lo.Contains(p.Icons, k) {
			p.Icons = append(p.Icons, k)
		}
	})
}

func (p *PageState) TitleString() string {
	if p.Title == "" {
		return util.AppName
	}
	if strings.HasPrefix(p.Title, "!") {
		return p.Title[1:]
	}
	return fmt.Sprintf("%s - %s", p.Title, util.AppName)
}

func (p *PageState) Username() string {
	if p.User != nil {
		return p.User.Name
	}
	return p.Profile.Name
}

func (p *PageState) AuthString() string {
	n := p.Profile.String()
	if p.User != nil {
		n = p.User.Name
	}
	msg := fmt.Sprintf("signed in as %s", n)
	if len(p.Accounts) == 0 {
		if n == user.DefaultProfile.Name {
			return "click to sign in"
		}
		return msg
	}
	return fmt.Sprintf("%s using [%s]", msg, p.Accounts.TitleString())
}

func (p *PageState) Clean(_ *http.Request, as *app.State) error {
	if p.Profile != nil && p.Profile.Theme == "" {
		p.Profile.Theme = theme.Default.Key
	}
	if p.RootIcon == "" {
		p.RootIcon = defaultIcon
	}
	if p.RootPath == "" {
		p.RootPath = "/"
	}
	if p.RootTitle == "" {
		p.RootTitle = defaultRootTitle
	}
	if defaultRootTitleAppend != "" {
		p.RootTitle += " " + defaultRootTitleAppend
	}
	if p.SearchPath == "" {
		p.SearchPath = DefaultSearchPath
	}
	if p.ProfilePath == "" {
		p.ProfilePath = DefaultProfilePath
	}
	if len(p.Menu) == 0 {
		m, data, err := cmenu.MenuFor(p.Context, p.Authed, p.Admin, p.Profile, p.Params, as, p.Logger)
		if err != nil {
			return err
		}
		if data != nil && p.Data == nil {
			p.Data = data
		}
		p.Menu = m
	}
	return nil
}

func (p *PageState) Close() {
	if p.Span != nil {
		p.Span.Complete()
	}
}

func (p *PageState) LogError(msg string, args ...any) {
	p.Logger.Errorf(msg, args...)
}

func (p *PageState) ClassDecl() string {
	if len(p.Icons) == 0 {
		return "-"
	}
	ret := &util.StringSlice{}
	if p.Profile.Mode != "" {
		ret.Push(p.Profile.ModeClass())
	}
	if p.Browser != "" {
		ret.Push("browser-" + p.Browser)
	}
	if p.OS != "" {
		ret.Push("os-" + p.OS)
	}
	if p.Platform != "" {
		ret.Push("platform-" + p.Platform)
	}
	if ret.Empty() {
		return ""
	}
	classes := ret.Join(" ")
	return fmt.Sprintf(` class=%q`, classes)
}

func (p *PageState) SetTitleAndData(title string, data any) {
	p.Title = title
	p.Data = data
}

func (p *PageState) MainClasses() string {
	var ret []string
	if p.HideHeader {
		ret = append(ret, "noheader")
	}
	if p.HideMenu {
		ret = append(ret, "nomenu")
	}
	return util.StringJoin(ret, " ")
}
