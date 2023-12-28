package domain

import (
	"time"
	"context"
)

type Mahasiswa struct {
	NIM			string 		`json:"nim"`
	Email		string 		`json:"email"`
	Password	string 		`json:"password"`
	Nama		string 		`json:"nama"`
	Jurusan		string		`json:"prodi"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	CreatedAt	time.Time 	`json:"created_at"`
}

type MahasiswaUsecase interface {
	FetchAll(ctx context.Context) ([]Mahasiswa, error)
	FetchByNIM(ctx context.Context, nim string) (Mahasiswa, error)
	Update(ctx context.Context, mhs *Mahasiswa) error 
	Register(ctx context.Context, mhs *Mahasiswa) error 
	Delete(ctx context.Context, nim string) error 
}

type MahasiswaRepostory interface {
	FetchAll(ctx context.Context) ([]Mahasiswa, error)
	FetchByNIM(ctx context.Context, nim string) (Mahasiswa, error)
	Update(ctx context.Context, mhs *Mahasiswa) error 
	Register(ctx context.Context, mhs *Mahasiswa) error 
	Delete(ctx context.Context, nim string) error 
}