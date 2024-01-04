package main

import (
	"fmt"
	"time"

	"github.com/devanfer02/game-corner-go/bootstrap/env"
	"github.com/devanfer02/game-corner-go/bootstrap/database/mysql"
	"github.com/devanfer02/game-corner-go/middleware"

	mhsrepomysql "github.com/devanfer02/game-corner-go/internal/mahasiswa/repository/mysql"
	mhsucase "github.com/devanfer02/game-corner-go/internal/mahasiswa/usecase"
	mhshandler "github.com/devanfer02/game-corner-go/internal/mahasiswa/delivery/http"

	"github.com/gin-gonic/gin"
)

func main() {
	env.InitEnv()
	db := mysql.NewMysqlDBConn()
	defer db.Close()

	router := gin.Default()	
	router.Use(middleware.CORS())
	//router.Use(middleware.InterceptorKey)

	//setting up mahasiswa layer
	mhsRepo := mhsrepomysql.NewMysqlMahasiswaRepository(db)
	mhsUcase := mhsucase.NewMahasiswaUsecase(mhsRepo, 10 * time.Second)
	mhshandler.NewMahasiswaHandler(router, mhsUcase)

	router.Run(fmt.Sprintf(":%s", env.Globenv.ServerAddress))
}