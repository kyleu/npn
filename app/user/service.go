package user

import "github.com/kyleu/npn/app/lib/filter"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("user", &filter.Ordering{Column: "created"})
}
