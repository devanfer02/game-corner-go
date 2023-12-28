package domain

import "context"

type LoginRequest struct {
	Email			string		`json:"email" validator:"required"`
	Password		string		`json:"password" validator:"required"`
}

type LoginResponse struct {
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
}

type LoginUsecase interface {
	FetchMahasiswaByEmail(ctx context.Context, email string) (Mahasiswa, error)
	CreateAcessToken(mahasiswa *Mahasiswa, secret string, expired int) (string, error)
	CreateRefreshToken(mahasiswa *Mahasiswa, secret string, expired int) (string, error)
}