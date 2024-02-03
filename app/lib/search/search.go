// Package search - Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/search/result"
	"github.com/kyleu/npn/app/lib/telemetry"
	"github.com/kyleu/npn/app/util"
)

type Provider func(context.Context, *Params, *app.State, *cutil.PageState, util.Logger) (result.Results, error)

func Search(ctx context.Context, params *Params, as *app.State, page *cutil.PageState) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", page.Logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	testFunc := func(ctx context.Context, p *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		return result.Results{{URL: "/search?q=test", Title: "Test Result", Icon: "star", Matches: nil}}, nil
	}
	allProviders = append(allProviders, testFunc)
	// $PF_SECTION_END(search_functions)$
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	params.Q = strings.TrimSpace(params.Q)

	results, errs := util.AsyncCollect(allProviders, func(p Provider) (result.Results, error) {
		return p(ctx, params, as, page, logger)
	})

	var ret result.Results = lo.FlatMap(results, func(x result.Results, _ int) []*result.Result {
		return x
	})
	return ret.Sort(), errs
}
