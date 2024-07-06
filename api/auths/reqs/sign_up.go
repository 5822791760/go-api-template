package reqs

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,min=1,max=255,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
	Name     string `json:"name" validate:"required,min=1,max=255"`
}
