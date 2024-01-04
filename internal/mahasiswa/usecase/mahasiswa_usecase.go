package usecase

import (
	"context"
	"log"
	"time"

	"github.com/devanfer02/game-corner-go/domain"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type mahasiswaUsecase struct {
	mahasiswaRepo  domain.MahasiswaRepository
	contextTimeout time.Duration
	mhsValidator *validator.Validate
}

func NewMahasiswaUsecase(repo domain.MahasiswaRepository, timeout time.Duration) domain.MahasiswaUsecase {
	return &mahasiswaUsecase{repo, timeout, validator.New()}
}

func (ucase *mahasiswaUsecase) FetchAll(ctx context.Context) ([]domain.Mahasiswa, error) {
	ctxc, cancel := context.WithTimeout(ctx, ucase.contextTimeout)
	defer cancel()

	res, err := ucase.mahasiswaRepo.FetchAll(ctxc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ucase *mahasiswaUsecase) FetchByNIM(ctx context.Context, nim string) (domain.Mahasiswa, error) {
	ctxc, cancel := context.WithTimeout(ctx, ucase.contextTimeout)
	defer cancel()

	res, err := ucase.mahasiswaRepo.FetchByNIM(ctxc, nim)
	if err != nil {
		return domain.Mahasiswa{}, err
	}

	return res, nil
}

func (ucase *mahasiswaUsecase) Register(ctx context.Context, mhs *domain.MahasiswaRegister) error {
	if err := ucase.mhsValidator.Struct(mhs); err != nil {
		return domain.ErrBadRequest
	}

	ctxc, cancel := context.WithTimeout(ctx, ucase.contextTimeout)
	defer cancel()

	existedMahasiswa, _ := ucase.FetchByNIM(ctxc, mhs.NIM)
	if existedMahasiswa != (domain.Mahasiswa{}) {
		return domain.ErrConflict
	}

	ucase.hashPassword(mhs)
	err := ucase.mahasiswaRepo.Register(ctxc, mhs)

	return err
}

func (ucase *mahasiswaUsecase) Update(ctx context.Context, mhs *domain.Mahasiswa) error {
	ctxc, cancel := context.WithTimeout(ctx, ucase.contextTimeout)
	defer cancel()

	mhs.UpdatedAt = time.Now()
	return ucase.mahasiswaRepo.Update(ctxc, mhs)
}

func (ucase *mahasiswaUsecase) Delete(ctx context.Context, nim string) error {
	ctxc, cancel := context.WithTimeout(ctx, ucase.contextTimeout)
	defer cancel()

	existedMahasiswa, err := ucase.mahasiswaRepo.FetchByNIM(ctxc, nim)
	if err != nil {
		return err
	}

	if existedMahasiswa == (domain.Mahasiswa{}) {
		return domain.ErrNotFound
	}

	return ucase.mahasiswaRepo.Delete(ctxc, nim)
}

func (ucase *mahasiswaUsecase) hashPassword(mhs *domain.MahasiswaRegister) {
	password, err := bcrypt.GenerateFromPassword([]byte(mhs.Password), 10)

	if err != nil {
		log.Fatalf("ERR: %s\n", err.Error())
	}

	mhs.Password = string(password)
}