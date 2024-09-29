package userusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Response =============================

type GetOneResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ============================== Usecase ==============================

func (u userUsecase) FindOne(ctx context.Context, id int) (GetOneResp, apperr.Err) {
	user, err := u.userRepo.FindOne(ctx, id)
	if err != nil {
		return GetOneResp{}, apperr.NewInternalServerErr(err)
	}

	if user == nil {
		return GetOneResp{}, apperr.NewUserNotFoundErr()
	}

	return GetOneResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
