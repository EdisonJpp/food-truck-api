package contract

type AuthenticateRequest struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"password" validate:"gte=3,required"`
}
