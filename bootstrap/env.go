package bootstrap

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	SecretKey      string
	Algorithm      string
	ExpiredMinutes int
	SystemTimeout  int
}

func NewEnv() *Env {
	expire, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRE_MINUTES"))
	if err != nil {
		log.Fatalf("Unable to read `ACCESS_TOKEN_EXPIRE_MINUTES` from env file.")
	}

	timeout, err := strconv.Atoi(os.Getenv("SYSTEM_TIMEOUT_SECOND"))
	if err != nil {
		log.Fatalf("Unable to read `SYSTEM_TIMEOUT_SECOND` from env file.")
	}

	return &Env{
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBName:         os.Getenv("DB_NAME"),
		SecretKey:      os.Getenv("SECRET_KEY"),
		Algorithm:      os.Getenv("ALGORITHM"),
		ExpiredMinutes: expire,
		SystemTimeout:  timeout,
	}
}
