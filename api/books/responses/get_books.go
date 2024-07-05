package responses

type author struct {
	ID   int64 `json:"id"`
	Name string `json:"name"`
}

type GetBooksResponse struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	BookNo string `json:"book_no"`

	Author author `json:"author"`
}
