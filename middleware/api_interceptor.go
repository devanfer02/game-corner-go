package middleware

import (
	"net/http"

	"github.com/devanfer02/game-corner-go/bootstrap/env"

	"github.com/gin-gonic/gin"
)

func InterceptorKey() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		key := ctx.Request.Header["X-API-Key"]
	
		if (key[0] == env.Globenv.ApiKey) {
			ctx.AbortWithStatus(http.StatusForbidden)		
			return
		}
	
		ctx.Next()
	}
}