package v1

import (
	"github/zine0/short_url/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterService(router *gin.RouterGroup) {
	v1 := router.Group("/v1")

	v1.POST("/add", service.Add)
}
