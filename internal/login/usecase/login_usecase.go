package usecase 

import (
	"context"
	"log"
	"time"

	"github.com/devanfer02/game-corner-go/domain"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	mahasiswaRepository domain.MahasiswaRepository
	contextTimeout time.Duration 
	validator validator.Validate
}

func (loginUcase *loginUsecase) FetchMahasiswaByNIM(ctx context.Context, email string) (domain.Mahasiswa, error) {

}

func (loginUcase *loginUsecase) CreateAcessToken(mahasiswa *domain.Mahasiswa, secret string, expired int) (string, error) {

}

func (loginUcase *loginUsecase) CreateRefreshToken(mahasiswa *domain.Mahasiswa, secret string, expired int) (string, error) {

}