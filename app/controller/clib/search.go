// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"net/http"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/lib/search"
	"github.com/kyleu/npn/views/vsearch"
)

const searchKey = "search"

func Search(w http.ResponseWriter, r *http.Request) {
	controller.Act(searchKey, w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := r.URL.Query().Get("q")
		params := &search.Params{Q: q, PS: ps.Params}
		results, errs := search.Search(ps.Context, params, as, ps)
		ps.SetTitleAndData("Search Results", results)
		if q != "" {
			ps.Title = fmt.Sprintf("[%s] %s", q, ps.Title)
		}
		if len(results) == 1 && results[0].URL != "" {
			return controller.FlashAndRedir(true, "single search result found", results[0].URL, ps)
		}
		ps.DefaultNavIcon = searchKey
		bc := []string{"Search||/search"}
		if q != "" {
			bc = append(bc, q)
		}
		return controller.Render(r, as, &vsearch.Results{Params: params, Results: results, Errors: errs}, ps, bc...)
	})
}
