package res

type books struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Bookno string `json:"book_no"`
}

type GetAuthorsResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`

	Books []books `json:"books"`
}
