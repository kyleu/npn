package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) CreateIfNeeded(ctx context.Context, userID uuid.UUID, username string, logger util.Logger) error {
	return nil
}
