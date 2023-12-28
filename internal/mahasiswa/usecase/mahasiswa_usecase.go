package usecase

import (
	"time"
	"context"

	"game-corner/domain"
)

type mahasiswaUsecase struct {
	mahasiswaRepo	domain.MahasiswaRepostory
	contextTimeout	time.Duration
}

func NewMahasiswaUsecase(repo domain.MahasiswaRepostory, timeout time.Duration) domain.MahasiswaUsecase {
	return &mahasiswaUsecase{repo, timeout}
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

func (ucase *mahasiswaUsecase) Register(ctx context.Context, mhs *domain.Mahasiswa) error {
	ctxc, cancel := context.WithTimeout(ctx, ucase.contextTimeout)
	defer cancel()

	existedMahasiswa, _ := ucase.FetchByNIM(ctxc, mhs.NIM)
	if existedMahasiswa != (domain.Mahasiswa{}) {
		return domain.ErrConflict
	}

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