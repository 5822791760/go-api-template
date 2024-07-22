package reqs

type CreateBook struct {
	Name     string  `json:"name" validate:"required,min=1,max=25"`
	Bookno   string  `json:"bookno" validate:"required,min=1,max=25"`
	Price    int32   `json:"price" validate:"required,gte=1"`
	Summary  *string `json:"summary"`
	Amount   int32   `json:"amount" validate:"required,gte=0"`
	AuthorID int32   `json:"author_id" validate:"required,gte=1"`
}
