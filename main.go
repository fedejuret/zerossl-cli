/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-cli/cmd"
	_ "github.com/fedejuret/zerossl-cli/cmd/modules/certificates"
	"github.com/fedejuret/zerossl-cli/lib/database"
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

	myFigure := figure.NewColorFigure("ZEROSSL CLI", "", "cyan", true)
	myFigure.Print()
	fmt.Println(color.YellowString(" By Federico Juretich <fedejuret@gmail.com>"))
	fmt.Println()
	fmt.Println()
	fmt.Println()

	if len(os.Getenv("ZEROSSL_API_KEY")) == 0 {
		fmt.Println(color.RedString("Environment variable ZEROSSL_API_KEY not found in your OS"))
		return
	}

	cmd.Execute()
}
