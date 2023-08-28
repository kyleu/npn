package controller

import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
)

// Initialize system dependencies for the marketing site.
func initSite(*app.State, util.Logger) {
}

// Configure marketing site data for each request.
func initSiteRequest(as *app.State, _ *cutil.PageState) error {
	return nil
}
