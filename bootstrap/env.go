package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DB_User     string `mapstructure:"DB_USER"`
	DB_Password string `mapstructure:"DB_PASSWORD"`
	DB_Host     string `mapstructure:"DB_HOST"`
	DB_Port     string `mapstructure:"DB_PORT"`
	DB_Name     string `mapstructure:"DB_NAME"`

	SecretKey      string `mapstructure:"SECRET_KEY"`
	ExpiredMinutes int    `mapstructure:"ACCESS_TOKEN_EXPIRE_MINUTES"`
	SystemTimeout  int    `mapstructure:"SYSTEM_TIMEOUT_SECOND"`

	BinanceAPI       string `mapstructure:"BINANCE_API_DOMAIN"`
	BinanceAccessKey string `mapstructure:"BINANCE_ACCESS_KEY"`
	BinanceSecretKey string `mapstructure:"BINANCE_SECRET_KEY"`

	OauthRedirectUri        string `mapstructure:"OAUTH_REDIRECT_URI"`
	GoogleOauthAuthUri      string `mapstructure:"GOOGLE_OAUTH_URI"`
	GoogleOauthTokenUri     string `mapstructure:"GOOGLE_OAUTH_TOKEN_URI"`
	GoogleOauthUserInfoUri  string `mapstructure:"GOOGLE_OAUTH_USERINFO_URI"`
	GoogleOauthClientID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleOauthClientSecret string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can not read file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Enviroment can not be loaded : ", err)
	}
	fmt.Println(env.GoogleOauthAuthUri)

	return &env
}
