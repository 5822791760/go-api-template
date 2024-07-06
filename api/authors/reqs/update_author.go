package reqs

type UpdateAuthorRequest struct {
	Name string `json:"name" validate:"min=1,max=255"`
	Bio  string `json:"bio" validate:"min=1,max=50"`
}
