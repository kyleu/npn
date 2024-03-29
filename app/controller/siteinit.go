package controller

import (
	"context"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/util"
)

// Initialize system dependencies for the marketing site.
func initSite(_ context.Context, _ *app.State, _ util.Logger) error {
	return nil
}

// Configure marketing site data for each request.
func initSiteRequest(as *app.State, _ *cutil.PageState) error {
	return nil
}
