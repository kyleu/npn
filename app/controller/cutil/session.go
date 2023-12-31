// Package cutil - Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"context"
	"slices"
	"strings"

	"github.com/mileusna/useragent"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/csession"
	"github.com/kyleu/npn/app/lib/telemetry"
	"github.com/kyleu/npn/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/npn/app/lib/user"
	"github.com/kyleu/npn/app/util"
)

var initialIcons = []string{"searchbox"}

func LoadPageState(as *app.State, rc *fasthttp.RequestCtx, key string, logger util.Logger) *PageState {
	parentCtx, logger := httpmetrics.ExtractHeaders(rc, logger)
	ctx, span, logger := telemetry.StartSpan(parentCtx, "http:"+key, logger)
	span.Attribute("path", string(rc.Request.URI().Path()))
	if !telemetry.SkipControllerMetrics {
		httpmetrics.InjectHTTP(rc, span)
	}
	session, flashes, prof, accts := loadSession(ctx, as, rc, logger)
	params := ParamSetFromRequest(rc)
	ua := useragent.Parse(string(rc.Request.Header.Peek("User-Agent")))
	os := strings.ToLower(ua.OS)
	browser := strings.ToLower(ua.Name)
	platform := "unknown"
	switch {
	case ua.Desktop:
		platform = "desktop"
	case ua.Tablet:
		platform = "tablet"
	case ua.Mobile:
		platform = "mobile"
	case ua.Bot:
		platform = "bot"
	}
	span.Attribute("browser", browser)
	span.Attribute("os", os)

	isAuthed, _ := user.Check("/", accts)
	isAdmin, _ := user.Check("/admin", accts)

	u, _ := as.User(ctx, prof.ID, logger)

	return &PageState{
		Action: key, Method: string(rc.Method()), URI: rc.Request.URI(), Flashes: flashes, Session: session,
		OS: os, OSVersion: ua.OSVersion, Browser: browser, BrowserVersion: ua.Version, Platform: platform,
		User: u, Profile: prof, Accounts: accts, Authed: isAuthed, Admin: isAdmin, Params: params,
		Icons: slices.Clone(initialIcons), Started: util.TimeCurrent(), Logger: logger, Context: ctx, Span: span,
	}
}

func loadSession(_ context.Context, _ *app.State, rc *fasthttp.RequestCtx, logger util.Logger) (util.ValueMap, []string, *user.Profile, user.Accounts) {
	sessionBytes := rc.Request.Header.Cookie(util.AppKey)
	session := util.ValueMap{}
	if len(sessionBytes) > 0 {
		dec, err := util.DecryptMessage(nil, string(sessionBytes), logger)
		if err != nil {
			logger.Warnf("error decrypting session: %+v", err)
		}
		err = util.FromJSON([]byte(dec), &session)
		if err != nil {
			session = util.ValueMap{}
		}
	}

	flashes := util.StringSplitAndTrim(session.GetStringOpt(csession.WebFlashKey), ";")
	if len(flashes) > 0 {
		delete(session, csession.WebFlashKey)
		err := csession.SaveSession(rc, session, logger)
		if err != nil {
			logger.Warnf("can't save session: %+v", err)
		}
	}

	prof, err := loadProfile(session)
	if err != nil {
		logger.Warnf("can't load profile: %+v", err)
	}

	var accts user.Accounts
	authX, ok := session[csession.WebAuthKey]
	if ok {
		authS, ok := authX.(string)
		if ok {
			accts = user.AccountsFromString(authS)
		}
	}

	if prof.ID == util.UUIDDefault {
		prof.ID = util.UUID()
		session["profile"] = prof
		err = csession.SaveSession(rc, session, logger)
		if err != nil {
			logger.Warnf("unable to save session for user [%s]", prof.ID.String())
			return nil, nil, prof, nil
		}
	}

	return session, flashes, prof, accts
}

func loadProfile(session util.ValueMap) (*user.Profile, error) {
	x, ok := session["profile"]
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	s, ok := x.(string)
	if !ok {
		m, ok := x.(map[string]any)
		if !ok {
			return user.DefaultProfile.Clone(), nil
		}
		s = util.ToJSON(m)
	}
	p := &user.Profile{}
	err := util.FromJSON([]byte(s), p)
	if err != nil {
		return nil, err
	}
	if p.Name == "" {
		p.Name = user.DefaultProfile.Name
	}
	return p, nil
}
