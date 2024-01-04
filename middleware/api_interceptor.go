package middleware

import (
	"net/http"

	"github.com/devanfer02/game-corner-go/bootstrap/env"

	"github.com/gin-gonic/gin"
)

func (mdlwr *Middleware) InterceptorKey() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		key := ctx.GetHeader("X-API-Key")
	
		if (key != env.Globenv.ApiKey) {
			ctx.AbortWithStatus(http.StatusForbidden)		
			return
		}
	
		ctx.Next()
	}
}