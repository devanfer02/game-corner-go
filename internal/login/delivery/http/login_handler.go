package http

import (
	"net/http"

	"github.com/devanfer02/game-corner-go/bootstrap/env"
	ghttp "github.com/devanfer02/game-corner-go/http"

	"github.com/devanfer02/game-corner-go/domain"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUsecase domain.LoginUsecase
	MahasiswaUcase domain.MahasiswaUsecase
}

func NewLoginHandler(r *gin.Engine, lucase domain.LoginUsecase, mucase domain.MahasiswaUsecase) {
	handler := LoginHandler{
		LoginUsecase: lucase,
		MahasiswaUcase: mucase,
	}

	r.POST("/mahasiswa/login", handler.LoginMahasiswa)
	r.POST("/mahasiswa/logout", handler.LogoutMahasiswa)
}

func (loginHandler *LoginHandler) LoginMahasiswa(ctx *gin.Context) {
	var request domain.LoginRequest 

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ghttp.ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	mhs, err := loginHandler.MahasiswaUcase.FetchByNIM(ctx.Request.Context(), request.NIM)
	if err != nil {
		if err == domain.ErrNotFound {
			err = domain.ErrInvalidMahasiswa
		}
		ctx.JSON(domain.GetErrorCode(err), ghttp.ResponseError{Status: domain.GetErrorCode(err), Message: err.Error()})
		return 
	}

	err = loginHandler.LoginUsecase.ValidateMahasiswa(&mhs, &request)	
	if err != nil {
		ctx.JSON(domain.GetErrorCode(err), ghttp.ResponseError{Status: domain.GetErrorCode(err), Message: err.Error()})
		return 
	}

	token, err := loginHandler.LoginUsecase.CreateFirebaseToken(&mhs, env.Globenv.AccessTokenExpyHour)
	ctx.Header("Authorization", token)
	ctx.JSON(http.StatusOK, ghttp.Response{Status: http.StatusOK, Message: "successfully login", Data: nil})
}

func (loginHanlder *LoginHandler) LogoutMahasiswa(ctx *gin.Context) {

}