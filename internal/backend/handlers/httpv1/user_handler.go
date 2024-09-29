package httpv1

import (
	"net/http"

	"github.com/5822791760/hr/internal/backend/usecases/userusecase"
	"github.com/5822791760/hr/pkg/apperr"
	"github.com/5822791760/hr/pkg/coreutil"
)

type UserHandler struct {
	db          coreutil.Transactionable
	userUsecase userusecase.UserUsecase
}

func NewAuthorHandler(db coreutil.Transactionable, userUsecase userusecase.UserUsecase) UserHandler {
	return UserHandler{
		db:          db,
		userUsecase: userUsecase,
	}
}

func (h UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	var err apperr.Err
	ctx := coreutil.GetContext(r, h.db)

	defer func() {
		coreutil.WriteError(w, err)
	}()

	id, err := coreutil.GetParamInt(r, "id")
	if err != nil {
		return
	}

	res, err := h.userUsecase.FindOne(ctx, id)
	if err != nil {
		return
	}

	coreutil.WriteJSON(w, http.StatusOK, res)
}
