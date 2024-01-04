package usecase

import (
	"context"
	"time"
	"fmt"

	"github.com/devanfer02/game-corner-go/domain"

	"firebase.google.com/go/v4/auth"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	mahasiswaRepository domain.MahasiswaRepository
	contextTimeout time.Duration 
	validator *validator.Validate
	fireAuth *auth.Client 
}

func NewLoginUsecase(mhsRepo domain.MahasiswaRepository, timeout time.Duration, auth *auth.Client) domain.LoginUsecase {
	return &loginUsecase{
		mahasiswaRepository: mhsRepo,
		contextTimeout: timeout,
		validator: validator.New(),
		fireAuth: auth,
	}
}

func (loginUcase *loginUsecase) ValidateMahasiswa(mahasiswa *domain.Mahasiswa, request *domain.LoginRequest) error {
	err := bcrypt.CompareHashAndPassword([]byte(mahasiswa.Password), []byte(request.Password))
	if err != nil {
		return domain.ErrInvalidMahasiswa
	}

	return nil 
}

func (loginUcase *loginUsecase) CreateFirebaseToken(mahasiswa *domain.Mahasiswa, expired int) (string, error) {
	expiresIn := time.Now().Add(time.Hour * time.Duration(expired))
	claims := map[string]interface{}{
		"nim": mahasiswa.NIM, 
		"expiresIn": expiresIn.Unix(),		
	}

	token, err := loginUcase.fireAuth.CustomTokenWithClaims(context.Background(), mahasiswa.NIM, claims)
	if err != nil {
		return "", fmt.Errorf("Error creating custom token. ERR: %s", err.Error())
	}

	return token, nil 
}