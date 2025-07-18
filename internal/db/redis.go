package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Rdb *redis.Client
var PgDb *pgx.Conn

func InitDB() {

	redisAddr := viper.GetStringMapString("redis")["addr"]
	redisPassword := viper.GetStringMapString("redis")["password"]

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
		PoolSize: 10,
	})

	Rdb = rdb

	pg := viper.GetStringMapString("pgsql")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", pg["username"], pg["password"], pg["addr"], pg["db"])
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic("error to connect pgsql")
	}

	PgDb = conn
}
