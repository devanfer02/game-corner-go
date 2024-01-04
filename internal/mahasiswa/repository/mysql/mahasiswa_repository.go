package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/devanfer02/game-corner-go/domain"
)

type mysqlMahasiswaRepostiory struct {
	Conn *sql.DB
}

func NewMysqlMahasiswaRepository(conn *sql.DB) domain.MahasiswaRepository {
	return &mysqlMahasiswaRepostiory{conn}
}

func (sql *mysqlMahasiswaRepostiory) fetch(ctx context.Context, query string, args ...any) ([]domain.Mahasiswa, error) {
	rows, err := sql.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		log.Printf("mahasiswa_repository.go ERR: %s\n", err.Error())
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("mahasiswa_repository.go ERR: %s\n", err.Error())
		}
	}()

	result := make([]domain.Mahasiswa, 0)
	for rows.Next() {
		mhs := domain.Mahasiswa{}
		err = rows.Scan(
			&mhs.NIM,
			&mhs.Email,
			&mhs.Password,
			&mhs.Nama,
			&mhs.Jurusan,
			&mhs.UpdatedAt,
			&mhs.CreatedAt,
		)

		if err != nil {
			log.Printf("mahasiswa_repository.go ERR: %s\n", err.Error())
		}

		result = append(result, mhs)
	}

	return result, nil

}

func (sql *mysqlMahasiswaRepostiory) FetchAll(ctx context.Context) ([]domain.Mahasiswa, error) {
	query := `SELECT nim, email, password, nama, jurusan, updated_at, created_at FROM mahasiswa`

	result, err := sql.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (sql *mysqlMahasiswaRepostiory) FetchByNIM(ctx context.Context, nim string) (domain.Mahasiswa, error) {
	query := `SELECT nim, email, password, nama, jurusan, updated_at, created_at FROM mahasiswa WHERE nim = ?`

	mhsList, err := sql.fetch(ctx, query, nim)

	if err != nil {
		return domain.Mahasiswa{}, err
	}

	if len(mhsList) == 0 {
		return domain.Mahasiswa{}, domain.ErrNotFound
	}

	return mhsList[0], nil
}

func (sql *mysqlMahasiswaRepostiory) Register(ctx context.Context, mhs *domain.MahasiswaRegister) error {
	query := `INSERT INTO mahasiswa (nim, email, password, nama, jurusan) 
				VALUES (?, ?, ?, ?, ?)`

	stmt, err := sql.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, mhs.NIM, mhs.Email, mhs.Password, mhs.Nama, mhs.Jurusan)
	if err != nil {
		return err
	}

	return nil
}

func (sql *mysqlMahasiswaRepostiory) Update(ctx context.Context, mhs *domain.Mahasiswa) error {
	query := `UPDATE mahasiswa SET nim = ?, email = ?, password = ?, nama = ?, jurusan = ?, updated_at = ?, created_at = ? WHERE nim = ?`

	stmt, err := sql.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, mhs.NIM, mhs.Email, mhs.Password, mhs.Nama, mhs.Jurusan, mhs.UpdatedAt, mhs.CreatedAt)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		return fmt.Errorf("Weird Behaviour. Total Affected: %d\n", affect)
	}

	return nil
}

func (sql *mysqlMahasiswaRepostiory) Delete(ctx context.Context, nim string) error {
	query := `DELETE FROM mahasiswa WHERE nim = ?`

	stmt, err := sql.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, nim)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		return fmt.Errorf("Weird Behaviour. Total Affected: %d\n", affect)
	}

	return nil
}
