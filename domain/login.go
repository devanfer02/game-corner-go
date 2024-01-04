package domain

import "context"

type LoginRequest struct {
	NIM				string		`json:"nim" validator:"required"`
	Password		string		`json:"password" validator:"required"`
}

type LoginResponse struct {
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
}

type LoginUsecase interface {
	FetchMahasiswaByNIM(ctx context.Context, nim string) (Mahasiswa, error)
	CreateAcessToken(mahasiswa *Mahasiswa, secret string, expired int) (string, error)
	CreateRefreshToken(mahasiswa *Mahasiswa, secret string, expired int) (string, error)
}