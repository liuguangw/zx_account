package request

type RegisterAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
