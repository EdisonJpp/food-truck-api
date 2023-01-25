package presenter

import "food-truck-api/package/entities"

type AuthenticatePresentation struct {
	client      *entities.Client
	accessToken *string
}

func AuthenticatePresent(token *string, client *entities.Client) AuthenticatePresentation {
	return AuthenticatePresentation{
		client:      client,
		accessToken: token,
	}
}
