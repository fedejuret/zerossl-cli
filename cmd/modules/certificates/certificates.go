/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	mainCmd "github.com/fedejuret/zerossl-cli/cmd"
	"github.com/spf13/cobra"
)

var certificatesCmd = &cobra.Command{
	Use:   "certificates",
	Short: "Interact with certificates",
	Long:  "Using this command, you can create, list, revoke platform certificates",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			cmd.Help()
		}
	},
}

func init() {
	mainCmd.RootCmd.AddCommand(certificatesCmd)
}
