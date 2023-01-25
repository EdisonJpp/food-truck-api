package presenter

type LoginPresentation struct {
	accessToken *string
}

func LoginPresent(token *string) LoginPresentation {
	return LoginPresentation{
		accessToken: token,
	}
}
