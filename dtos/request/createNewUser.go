package request

type CreateNewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
