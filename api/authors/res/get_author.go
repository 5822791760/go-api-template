package res

type getAuthorBooks struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Bookno string `json:"book_no"`
}

type GetAuthorResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`

	Books []getAuthorBooks `json:"books"`
}
