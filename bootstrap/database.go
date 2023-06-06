package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Databases struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewDatabases(env *Env) *Databases {
	db, err := NewPostgresDatabase(env)
	if err != nil {
		log.Fatal(err)
	}
	return &Databases{
		DB:  db,
		RDB: NewRedisClient(env),
	}
}

func NewPostgresDatabase(env *Env) (*gorm.DB, error) {

	username := env.DB_User
	password := env.DB_Password
	host := env.DB_Host
	port := env.DB_Port
	dbname := env.DB_Name

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	return db, err
}

func NewRedisClient(env *Env) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.Redis_Host, env.Redis_Port),
		Password: env.Redis_Password,
		DB:       env.Redis_DB,
		PoolSize: env.Redis_PoolSize,
	})
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatal("Can not connect to redis, got ", err)
	}
	return rdb
}
