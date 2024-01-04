package main

import (
	"fmt"
	"time"

	"github.com/devanfer02/game-corner-go/bootstrap/database/mysql"
	"github.com/devanfer02/game-corner-go/bootstrap/env"
	"github.com/devanfer02/game-corner-go/bootstrap/storage/firebase"
	"github.com/devanfer02/game-corner-go/middleware"

	mhshandler "github.com/devanfer02/game-corner-go/internal/mahasiswa/delivery/http"
	mhsrepomysql "github.com/devanfer02/game-corner-go/internal/mahasiswa/repository/mysql"
	mhsucase "github.com/devanfer02/game-corner-go/internal/mahasiswa/usecase"

	authucase "github.com/devanfer02/game-corner-go/internal/login/usecase"
	authhandler "github.com/devanfer02/game-corner-go/internal/login/delivery/http"

	"github.com/gin-gonic/gin"
)

func main() {
	env.InitEnv()
	db := mysql.NewMysqlDBConn()
	defer db.Close()

	authClient := firebase.CreateAuthClient()
	mdlwr := middleware.NewMiddleware(authClient)

	router := gin.Default()	
	router.Use(mdlwr.CORS())
	router.Use(mdlwr.InterceptorKey())

	//setting up mahasiswa layer
	mhsRepo := mhsrepomysql.NewMysqlMahasiswaRepository(db)
	mhsUcase := mhsucase.NewMahasiswaUsecase(mhsRepo, 10 * time.Second)
	mhshandler.NewMahasiswaHandler(router, mhsUcase, mdlwr)

	//setting up login layer
	loginUcase := authucase.NewLoginUsecase(mhsRepo, 10 * time.Second, authClient)
	authhandler.NewLoginHandler(router, loginUcase, mhsUcase)

	router.Run(fmt.Sprintf(":%s", env.Globenv.ServerAddress))
}