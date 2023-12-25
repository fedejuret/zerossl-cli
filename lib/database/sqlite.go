package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Database() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", os.Getenv("ZEROSSL_FOLDER")+"/database.db")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InitializeDatabase() {

	db, err := Database()

	if err != nil {
		log.Fatal(err)
	}

	// Create certificates table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS certificates (
			id integer not null primary key AUTOINCREMENT,
			hash string not null unique,
			cname text not null,
			validation_method integer not null
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
