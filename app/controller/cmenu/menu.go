package cmenu

import (
	"context"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/lib/menu"
	"github.com/kyleu/npn/app/lib/telemetry"
	"github.com/kyleu/npn/app/lib/user"
	"github.com/kyleu/npn/app/util"
)

func MenuFor(
	ctx context.Context, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, as *app.State, logger util.Logger,
) (menu.Items, any, error) {
	ctx, sp, _ := telemetry.StartSpan(ctx, "menu", logger)
	defer sp.Complete()
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
	ws := &menu.Item{Key: "workspace", Title: "Workspace", Description: "A work-in-progress", Icon: "star", Route: "/x"}
	ret = append(ret, ws, menu.Separator)
	// This is your menu, feel free to customize it
	admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
	ret = append(ret, admin)
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
