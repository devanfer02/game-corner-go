package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/devanfer02/game-corner-go/bootstrap/env"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDBConn() *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.Globenv.DBUser, 
		env.Globenv.DBPassword,
		env.Globenv.DBHost, 
		env.Globenv.DBPort,
		env.Globenv.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not open database. ERR: %s\n", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping database. ERR: %s\n", err.Error())
	}

	migrate(
		db, 
		"bootstrap/database/mysql/migrations/create_mahasiswa_table.sql",
	)

	return db 
}

func migrate(db *sql.DB, filepaths ...string) {
	for _, filename := range filepaths {
		sqlfile, err := os.ReadFile(filename)

		if err != nil {
			log.Fatalf("Failed to run migration. ERR:%s\n", err.Error())
		}

		_, err = db.Exec(string(sqlfile))

		if err != nil {
			log.Fatalf("Failed to run sql query. ERR:%s\n", err.Error())
		}

		log.Printf("migration file %s success\n", filename)
	}
}