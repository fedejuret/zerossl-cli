/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/api/structs/requests"
	certificate_service "github.com/fedejuret/zerossl-golang-cli/lib/services"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:     "verify",
	Aliases: []string{"validate"},
	Short:   "Verify certificates",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		certId := args[0]

		cname, verificationMethod, err := certificate_service.GetByHash(certId)

		if err != nil {
			log.Fatal(err)
		}

		if cname == "not found" { // must ask for verification method
			fmt.Println(color.RedString("That certificate was not generated using this tool or was not found in sqlite database"))
		} else {

			var verificationRequest requests.VerifyDomainStructure
			if verificationMethod == 1 { // file upload
				verificationRequest = requests.VerifyDomainStructure{
					ValidationMethod:  "HTTP_CSR_HASH",
					VerificationEmail: "",
				}
			}

			response := api.Post("/certificates/"+certId+"/challenges", verificationRequest)

			fmt.Println(string(response))

		}
	},
}

func init() {
	certificatesCmd.AddCommand(verifyCmd)

}
