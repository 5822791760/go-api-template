package res

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	LastSignInAt string `json:"last_sign_in_at"`
}
