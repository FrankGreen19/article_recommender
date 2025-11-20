package output

type LoginOutputDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginOutputDto(access string, refresh string) LoginOutputDto {
	return LoginOutputDto{
		access,
		refresh,
	}
}
