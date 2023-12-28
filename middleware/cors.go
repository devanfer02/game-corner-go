package middleware 

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"gitub.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: 		[]string{os.Getenv("CLIENT_URL")},
		AllowMethods:		[]string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders: 		[]string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders: 		[]string{"Content-Length"},	
		AllowCredentials: 	true,	
	})
}