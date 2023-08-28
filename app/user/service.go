package user

import (
	"path"

	"github.com/google/uuid"

	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/util"
)

type Service struct {
	files  filesystem.FileLoader
	logger util.Logger
}

func NewService(f filesystem.FileLoader, logger util.Logger) *Service {
	return &Service{files: f, logger: logger}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("user", &filter.Ordering{Column: "created"})
}

func dirFor(userID uuid.UUID) string {
	return path.Join("users", userID.String())
}
