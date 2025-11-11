package dto

type JwtPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJwtPair(accessToken string, refreshToken string) JwtPair {
	return JwtPair{AccessToken: accessToken, RefreshToken: refreshToken}
}
