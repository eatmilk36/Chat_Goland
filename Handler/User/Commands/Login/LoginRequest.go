package Login

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
