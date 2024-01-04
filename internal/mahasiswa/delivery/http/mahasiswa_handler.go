package handler

import (
	"net/http"

	"github.com/devanfer02/game-corner-go/domain"
	ghttp "github.com/devanfer02/game-corner-go/http"

	"github.com/gin-gonic/gin"
)

type MahasiswaHandler struct {
	MahasiswaUsecase domain.MahasiswaUsecase
}

func NewMahasiswaHandler(r *gin.Engine, u domain.MahasiswaUsecase) {
	handler := &MahasiswaHandler{
		MahasiswaUsecase: u,
	}

	rMhs := r.Group("/mahasiswa")
	{
		rMhs.GET("/", handler.FetchMahasiswa)
		rMhs.GET("/:nim", handler.FetchByNIM)
		rMhs.POST("/", handler.RegisterMahasiswa)
		rMhs.PATCH("/:nim", handler.UpdateMahasiswa)
		rMhs.DELETE("/:nim", handler.DeleteMahasiswa)
	}
}

func (handler *MahasiswaHandler) FetchMahasiswa(ctx *gin.Context) {
	listMhs, err := handler.MahasiswaUsecase.FetchAll(ctx.Request.Context())

	code := domain.GetErrorCode(err)
	if err != nil {
		ctx.JSON(code, ghttp.ResponseError{Status: code, Message: err.Error()})
		return
	}

	ctx.JSON(code, ghttp.Response{Status: code, Message: "successfully fetch all mahasiswa", Data: listMhs})
}

func (handler *MahasiswaHandler) FetchByNIM(ctx *gin.Context) {
	nimParam := ctx.Param("nim")
	if nimParam == "" {
		ctx.JSON(http.StatusNotFound, ghttp.ResponseError{Status: http.StatusNotFound, Message: domain.ErrNotFound.Error()})
		return
	}

	mhs, err := handler.MahasiswaUsecase.FetchByNIM(ctx.Request.Context(), nimParam)
	code := domain.GetErrorCode(err)
	if err != nil {
		ctx.JSON(code, ghttp.ResponseError{Status: code, Message: err.Error()})
		return
	}

	ctx.JSON(code, ghttp.Response{Status: code, Message: "successfully fetch mahasiswa", Data: mhs})
}

func (handler *MahasiswaHandler) RegisterMahasiswa(ctx *gin.Context) {
	var mahasiswa domain.MahasiswaRegister

	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		ctx.JSON(http.StatusBadRequest, ghttp.ResponseError{Status: http.StatusBadRequest, Message: err.Error()})
		return 
	}	

	err := handler.MahasiswaUsecase.Register(ctx.Request.Context(), &mahasiswa)
	code := domain.GetErrorCode(err)

	if err != nil {
		ctx.JSON(code, ghttp.Response{Status: code, Message: err.Error()})
		return 
	}

	ctx.JSON(code, ghttp.Response{Status: code, Message:"successfully register mahasiswa" ,Data: mahasiswa})
}

func (handler *MahasiswaHandler) UpdateMahasiswa(ctx *gin.Context) {

}

func (handler *MahasiswaHandler) DeleteMahasiswa(ctx *gin.Context) {
	
}