package user

import (
	"Rest-Api-learning/pkg/logging"
	"context"
)

type Service struct {
	storage Storage
	Logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO for next one
	return
}
