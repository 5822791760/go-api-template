package res

type UpdateBook struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name"`
	Bookno    string   `json:"bookno"`
	Price     *float64 `json:"price"`
	Summary   *string  `json:"summary"`
	AuthorID  *int32   `json:"author_id"`
	Amount    int32    `json:"amount"`
	UpdatedAt string   `json:"updated_at"`
}
