package user

import (
	"context"
	"github.com/kyleu/npn/app/util"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, logger util.Logger, models ...*User) error {
	return nil
}

func (s *Service) Update(ctx context.Context, _ any, model *User, logger util.Logger) error {
	return nil
}

func (s *Service) Save(ctx context.Context, logger util.Logger, models ...*User) error {
	return nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID, logger util.Logger) error {
	return nil
}
