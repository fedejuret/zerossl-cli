package certificate_service

import (
	"log"

	"github.com/fedejuret/zerossl-cli/lib/database"
	"github.com/fedejuret/zerossl-cli/lib/models"
)

func Store(certificate models.Certificate, validationMethod int8) {

	db, _ := database.Database()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into certificates(hash, cname, validation_method) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(certificate.ID, certificate.CommonName, validationMethod)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}

func GetByHash(hash string) (cname string, verificationMethod int, err error) {

	db, _ := database.Database()

	stmt, err := db.Prepare("SELECT cname, validation_method FROM certificates WHERE hash = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(hash).Scan(&cname, &verificationMethod)

	if err != nil {
		return "not found", -1, nil
	}

	defer db.Close()

	return cname, verificationMethod, nil

}
