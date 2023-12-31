package domain

import (
	"time"
	"context"
)

type MahasiswaRegister struct {
	NIM			string 		`json:"nim" validate:"required"`
	Email		string 		`json:"email" validate:"required,email"`
	Password	string 		`json:"password" validate:"required,alphanum,min=6,max=30"`
	Nama		string 		`json:"nama" validate:"required"`
	Jurusan		string		`json:"prodi" validate:"required"`
}

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
	Register(ctx context.Context, mhs *MahasiswaRegister) error 
	Delete(ctx context.Context, nim string) error 
}

type MahasiswaRepository interface {
	FetchAll(ctx context.Context) ([]Mahasiswa, error)
	FetchByNIM(ctx context.Context, nim string) (Mahasiswa, error)
	Update(ctx context.Context, mhs *Mahasiswa) error 
	Register(ctx context.Context, mhs *MahasiswaRegister) error 
	Delete(ctx context.Context, nim string) error 
}