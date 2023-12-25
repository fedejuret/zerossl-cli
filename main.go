/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package main

import (
	"github.com/fedejuret/zerossl-golang-cli/cmd"
	_ "github.com/fedejuret/zerossl-golang-cli/cmd/modules/certificates"
	"github.com/fedejuret/zerossl-golang-cli/lib/database"
)

func main() {

	database.InitializeDatabase()

	cmd.Execute()
}
