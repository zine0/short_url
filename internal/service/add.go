package service

import (
	"context"
	"fmt"
	"github/zine0/short_url/internal/db"
	"github/zine0/short_url/internal/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spf13/viper"
)

type message struct {
	Url string
}

func Add(ctx *gin.Context) {
	msg := message{}
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error bind json",
		})
		return
	}

	msg.Url = utils.TrimUrl(msg.Url)

	nextId, err := nextId()
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}

	key := utils.IntToBase62(nextId)

	ctxC, cancel := context.WithTimeout(ctx, 500*time.Microsecond)
	defer cancel()

	db.Rdb.Set(ctxC, key, msg.Url, time.Hour*8)
	query := db.New(db.PgDb)

	err = query.AddUrl(context.Background(), db.AddUrlParams{
		Key: key,
		Url: pgtype.Text{String: msg.Url, Valid: true},
	})

	if err != nil {
		fmt.Printf("error to insert into pgsql: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}

	var builder strings.Builder
	builder.WriteString("http://")
	builder.WriteString(viper.GetStringMapString("app")["host"])
	builder.WriteRune('/')
	builder.WriteString(key)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":      "success",
		"shortUrl": builder.String(),
	})

}

func nextId() (int64, error) {
	return db.Rdb.Incr(context.Background(), "url:id").Result()
}
