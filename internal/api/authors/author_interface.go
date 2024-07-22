package authors

import (
	"github.com/5822791760/go-api-template/internal/api/authors/reqs"
	"github.com/5822791760/go-api-template/internal/api/authors/res"
	"github.com/5822791760/go-api-template/pkg/errs"
)

type IAuthorUseCase interface {
	GetAuthors() ([]res.GetAuthors, errs.ErrRenderer)
	GetAuthor(id int32) (res.GetAuthor, errs.ErrRenderer)
	CreateAuthor(body reqs.CreateAuthor) (res.CreateAuthor, errs.ErrRenderer)
	UpdateAuthor(id int32, body reqs.UpdateAuthor) (res.UpdateAuthor, errs.ErrRenderer)
}
