package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mdlwr *Middleware) ValidateFirebaseToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")

		token, err := mdlwr.fireAuth.VerifyIDToken(context.Background(), headerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		ctx.Set("mahasiswa", token)
		ctx.Next()
	}
}