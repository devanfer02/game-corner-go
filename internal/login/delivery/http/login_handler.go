package http

import (
	"github.com/devanfer02/game-corner-go/domain"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUsecase domain.LoginUsecase
}

func NewLoginHandler(r *gin.Engine, ucase domain.LoginUsecase) {
	handler := LoginHandler{LoginUsecase: ucase}

	r.POST("/mahasiswa/login", handler.LoginMahasiswa)
	r.POST("/mahasiswa/logout", handler.LogoutMahasiswa)
}

func (loginHandler *LoginHandler) LoginMahasiswa(ctx *gin.Context) {

}

func (loginHanlder *LoginHandler) LogoutMahasiswa(ctx *gin.Context) {

}