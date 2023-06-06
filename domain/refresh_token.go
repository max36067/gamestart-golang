package domain

type RefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenUsecase interface {
	GetUserByID(id int) (User, error)
	CreateAccessToken(user *User, secret string, expire int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expire int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (int, error)
}
