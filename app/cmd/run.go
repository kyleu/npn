package cmd

import (
	"context"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/util"
)

func Run(ctx context.Context, bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd(ctx).Execute(); err != nil {
		return util.RootLogger, err
	}
	return util.RootLogger, nil
}
