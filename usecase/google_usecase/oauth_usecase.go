package google_usecase

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type googleOauthUsecase struct {
	env *bootstrap.Env
}

func NewGoogleOauthUsecase(env *bootstrap.Env) domain.GoogleOauthUsecase {
	return &googleOauthUsecase{
		env: env,
	}
}

func (gou *googleOauthUsecase) RequestAccessToken(code string) (domain.GoogleOauthTokenResponse, error) {
	var token domain.GoogleOauthTokenResponse
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", gou.env.GoogleOauthClientID)
	data.Set("client_secret", gou.env.GoogleOauthClientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", gou.env.OauthRedirectUri)

	resp, err := http.Post(gou.env.GoogleOauthTokenUri, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return token, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()
	err = json.Unmarshal(body, &token)
	if err != nil {
		return token, err
	}
	fmt.Println(token.AccessToken)
	return token, nil
}

func (gou *googleOauthUsecase) GetUserInfo(googleOauthTokenResponse *domain.GoogleOauthTokenResponse) (domain.GoogleUser, error) {
	var userInfo domain.GoogleUser

	queryParam := url.Values{}
	queryParam.Set("alt", "json")
	// queryParam.Set("access_token", googleOauthTokenResponse.AccessToken)
	request, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", gou.env.GoogleOauthUserInfoUri, queryParam.Encode()), nil)
	if err != nil {
		log.Println("Making request failed")
		return userInfo, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", googleOauthTokenResponse.TokenType, googleOauthTokenResponse.AccessToken))

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println("Request failed from google user info api.")
		return userInfo, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read body failed.")
		return userInfo, err
	}
	log.Println(string(body))
	defer resp.Body.Close()
	err = json.Unmarshal([]byte(body), &userInfo)
	if err != nil {
		log.Println("JSON unmarshal failed.")
		return userInfo, err
	}

	return userInfo, nil
}
