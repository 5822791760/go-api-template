package userusecase

import (
	"context"

	"github.com/5822791760/hr/internal/backend/repos"
	"github.com/5822791760/hr/pkg/apperr"
	"github.com/samber/lo"
)

// ============================== Response =============================

type FindAllResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ============================== Usecase ==============================

func (u userUsecase) FindAll(ctx context.Context) ([]FindAllResp, apperr.Err) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return []FindAllResp{}, apperr.NewInternalServerErr(err)
	}

	return lo.Map(users, func(user repos.User, _ int) FindAllResp {
		return FindAllResp{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}
	}), nil
}
