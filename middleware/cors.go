package middleware 

import (
	"net/http"

	"github.com/devanfer02/game-corner-go/bootstrap/env"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: 		[]string{env.Globenv.ClientURL},
		AllowMethods:		[]string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders: 		[]string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders: 		[]string{"Content-Length"},	
		AllowCredentials: 	true,	
	})
}