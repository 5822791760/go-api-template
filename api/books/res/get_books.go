package res

type author struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type GetBooks struct {
	ID      int32   `json:"id"`
	Name    string  `json:"name"`
	BookNo  string  `json:"book_no"`
	Summary *string `json:"summary"`

	Author *author `json:"author"`
}
