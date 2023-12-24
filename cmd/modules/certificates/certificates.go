/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"fmt"

	mainCmd "github.com/fedejuret/zerossl-golang-cli/cmd"
	"github.com/spf13/cobra"
)

// certificatesCmd represents the create command
var certificatesCmd = &cobra.Command{
	Use:   "certificates",
	Short: "Interact with certificates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	mainCmd.RootCmd.AddCommand(certificatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// certificatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// certificatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
