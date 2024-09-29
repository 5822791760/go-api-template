package userusecase

import (
	"context"

	"github.com/5822791760/hr/internal/backend/repos"
	"github.com/5822791760/hr/pkg/apperr"
)

type userUsecase struct {
	userRepo repos.UserRepo
}

func NewUserUsecase(userRepo repos.UserRepo) userUsecase {
	return userUsecase{
		userRepo: userRepo,
	}
}

type UserUsecase interface {
	FindOne(ctx context.Context, id int) (GetOneResp, apperr.Err)
}
