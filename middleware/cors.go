package middleware 

import (
	"github.com/gin-gonic/gin"
)

func (mdlwr *Middleware) CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {		
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5600")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-Key",)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return 
		}

		ctx.Next()
	}
}