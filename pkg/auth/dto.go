package auth

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `jsno:"password" validate:"required,min=5" `
}
type GoogleResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Verified    bool   `json:"verified_email"`
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Given_name  string `json:"given_name"`
	Family_name string `json:"family_name"`
}
