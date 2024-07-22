package auths

type SignInToken struct {
	AccessToken  string
	LastSignInAt string
}

type SignUpBody struct {
	Email    string
	Password string
	Name     string
}
