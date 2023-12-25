/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fedejuret/zerossl-golang-cli/cmd"
	_ "github.com/fedejuret/zerossl-golang-cli/cmd/modules/certificates"
	"github.com/fedejuret/zerossl-golang-cli/lib/database"
)

var ProjectFolder string = os.Getenv("HOME") + "/.zerossl-cli/"

func main() {

	os.Setenv("ZEROSSL_FOLDER", ProjectFolder)

	if _, err := os.Stat(ProjectFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(ProjectFolder, 0700)

		if err != nil {
			fmt.Println("Error when try to create '" + ProjectFolder + "' folder:")
			log.Fatal(err)
		}
	}
	database.InitializeDatabase()

	cmd.Execute()
}
