package web

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
