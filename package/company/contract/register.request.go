package contract

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"gte=3,required"`
}
