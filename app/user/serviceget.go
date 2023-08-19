package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/npn/app/lib/filter"
	"github.com/kyleu/npn/app/util"
)

func (s *Service) List(ctx context.Context, params *filter.Params, logger util.Logger) (Users, error) {
	return nil, nil
}

func (s *Service) Count(ctx context.Context, whereClause string, logger util.Logger, args ...any) (int, error) {
	return 0, nil
}

func (s *Service) Get(ctx context.Context, _ any, id uuid.UUID, logger util.Logger) (*User, error) {
	return nil, nil
}

func (s *Service) GetMultiple(ctx context.Context, logger util.Logger, ids ...uuid.UUID) (Users, error) {
	return nil, nil
}

func (s *Service) Search(ctx context.Context, query string, params *filter.Params, logger util.Logger) (Users, error) {
	return nil, nil
}
