package reqs

type UpdateAuthor struct {
	Name string `json:"name" validate:"min=1,max=255"`
	Bio  string `json:"bio" validate:"min=1,max=50"`
}
