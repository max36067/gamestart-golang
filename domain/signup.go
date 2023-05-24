package domain

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupUsecase interface {
	CreateUser(user *User) error
	CreateSalt(salt *Salt) error
	GetUserByEmail(email string) (User, error)
	CreateAccessToken(user *User, secret string, expire int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expire int) (refreshToken string, err error)
}
