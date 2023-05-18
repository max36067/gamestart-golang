package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Database struct {
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		Name     string `mapstructure:"DB_NAME"`
	}
	Server struct {
		SecretKey      string `mapstructure:"SECRET_KEY"`
		ExpiredMinutes int    `mapstructure:"ACCESS_TOKEN_EXPIRE_MINUTES"`
		SystemTimeout  int    `mapstructure:"SYSTEM_TIMEOUT_SECOND"`
	}
	Binance struct {
		API       string `mapstructure:"BINANCE_API_DOMAIN"`
		AccessKey string `mapstructure:"BINANCE_ACCESS_KEY"`
		SecretKey string `mapstructure:"BINANCE_SECRET_KEY"`
	}
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

	return &env
}
