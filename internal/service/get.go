package service

import (
	"context"
	"errors"
	"github/zine0/short_url/internal/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Get(ctx *gin.Context) {
	key := ctx.Param("key")
	ctxC, cancle := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancle()
	val, err := db.Rdb.Get(ctxC, key).Result()

	if err != nil && !errors.Is(err, redis.Nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}

	if err != nil && errors.Is(err, redis.Nil) {
		query := db.New(db.PgDb)
		ctxCC, cancle_ := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancle_()
		url, err := query.GetUrl(ctxCC, key)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "error",
			})
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, url.Url.String)
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, val)
}
