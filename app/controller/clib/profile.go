// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller"
	"github.com/kyleu/npn/app/controller/csession"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/theme"
	"github.com/kyleu/npn/app/util"
	"github.com/kyleu/npn/views/vprofile"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	controller.Act("profile", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(w, r, as, ps)
	})
}

func ProfileSite(w http.ResponseWriter, r *http.Request) {
	controller.ActSite("profile", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(w, r, as, ps)
	})
}

func profileAction(w http.ResponseWriter, r *http.Request, as *app.State, ps *cutil.PageState) (string, error) {
	ps.SetTitleAndData("Profile", ps.Profile)
	thm := as.Themes.Get(ps.Profile.Theme, ps.Logger)

	prvs, err := as.Auth.Providers(ps.Logger)
	if err != nil {
		return "", errors.Wrap(err, "can't load providers")
	}

	redir := "/"
	ref := r.Header.Get("Referer")
	if ref != "" {
		u, err := url.Parse(ref)
		if err == nil && u != nil && u.Path != cutil.DefaultProfilePath {
			redir = u.Path
		}
	}
	ps.DefaultNavIcon = "profile"
	page := &vprofile.Profile{Profile: ps.Profile, Theme: thm, Providers: prvs, Referrer: redir}
	return controller.Render(r, as, page, ps, "Profile")
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	controller.Act("profile.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}

		n := ps.Profile.Clone()

		referrerDefault := frm.GetStringOpt("referrer")
		if referrerDefault == "" {
			referrerDefault = cutil.DefaultProfilePath
		}

		n.Name = frm.GetStringOpt("name")
		n.Mode = frm.GetStringOpt("mode")
		n.Theme = frm.GetStringOpt("theme")
		if n.Theme == theme.Default.Key {
			n.Theme = ""
		}
		if ps.Profile.ID == util.UUIDDefault {
			n.ID = util.UUID()
		} else {
			n.ID = ps.Profile.ID
		}

		err = csession.SaveProfile(n, w, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		curr, _ := as.Services.User.Get(ps.Context, nil, ps.Profile.ID, ps.Logger)
		if curr != nil {
			curr.Name = n.Name
			if curr.Picture == "" {
				curr.Picture = ps.Accounts.Image()
			}
			err = as.Services.User.Update(ps.Context, nil, curr, ps.Logger)
			if err != nil {
				return "", err
			}
		}

		return controller.ReturnToReferrer("Saved profile", referrerDefault, ps)
	})
}
