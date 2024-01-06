package auth

type LoginInput struct {
	Email    string `json:"Email" validate:"required"`
	Password string `jsno:"password" validate:"required,min=8" `
}
