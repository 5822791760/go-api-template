package userusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Response =============================

type FindOneResp struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ============================== Usecase ==============================

func (u userUsecase) FindOne(ctx context.Context, id int64) (FindOneResp, apperr.Err) {
	user, err := u.userRepo.FindOne(ctx, id)
	if err != nil {
		return FindOneResp{}, apperr.NewInternalServerErr(err)
	}

	if user == nil {
		return FindOneResp{}, apperr.NewUserNotFoundErr()
	}

	return FindOneResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
