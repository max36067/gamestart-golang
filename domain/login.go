package domain

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	GetUserByEmail(email string) (User, error)
	CreateAccessToken(user *User, secret string, expire int) (access_token string, err error)
	CreateRefreshToken(user *User, secret string, expire int) (refreshToken string, err error)
	VerifyPassword(hashedPassword, candidatePassword string) error
}
