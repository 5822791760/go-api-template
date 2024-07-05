package requests

type CreateAuthorRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	Bio  string `json:"bio" validate:"required,min=1,max=50"`
}
