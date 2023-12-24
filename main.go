/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/fedejuret/zerossl-golang-cli/cmd"
	_ "github.com/fedejuret/zerossl-golang-cli/cmd/modules/certificates"
)

func main() {
	cmd.Execute()
}
