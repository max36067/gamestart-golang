package domain

type GoogleOauthTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpireIn    int    `json:"expires_in,omitempty"`
	IdToken     string `json:"id_token,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

type GoogleOauthUsecase interface {
	RequestAccessToken(code string) (GoogleOauthTokenResponse, error)
	GetUserInfo(googleOauthTokenResponse *GoogleOauthTokenResponse) (GoogleUser, error)
}
