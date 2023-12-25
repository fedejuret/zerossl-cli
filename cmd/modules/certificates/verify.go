/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/api/structs/requests"
	"github.com/fedejuret/zerossl-golang-cli/lib/models"
	certificate_service "github.com/fedejuret/zerossl-golang-cli/lib/services"
	"github.com/fedejuret/zerossl-golang-cli/lib/utils"
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
			} else if verificationMethod == 2 { // dns add cname record
				verificationRequest = requests.VerifyDomainStructure{
					ValidationMethod:  "CNAME_CSR_HASH",
					VerificationEmail: "",
				}
			}

			spinner := utils.GetSpinner("Executing validation...", "fgYellow")
			spinner.Start()
			response := api.Post("/certificates/"+certId+"/challenges", verificationRequest)
			spinner.Stop()

			var responseJson map[string]interface{}

			if err := json.Unmarshal(response, &responseJson); err != nil {
				log.Fatal(err)
			}

			success, successExists := responseJson["success"]
			if successExists {
				if reflect.TypeOf(success).Kind() == reflect.Bool && success == false {
					errorInfo, errorExists := responseJson["error"].(map[string]interface{})
					if errorExists {
						errorMessage, typeExists := errorInfo["type"].(string)
						if typeExists {
							fmt.Println(color.RedString("Error when trying to verify domain: " + errorMessage))
						} else {
							fmt.Println(color.RedString("Error when trying to verify domain: unknown error type"))
						}
					} else {
						fmt.Println(color.RedString("Error when trying to verify domain: unknown error information"))
					}

				}
			} else {

				certificate, err := models.UnmarshalCertificate(response)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(color.GreenString("SUCCESS! The domain"), certificate.CommonName, color.GreenString("was validated successfully"))
				fmt.Println(color.HiMagentaString("Now you must wait a few minutes until the entity issues the certificate to be able to download it and upload it to your site"))

			}

		}
	},
}

func init() {
	certificatesCmd.AddCommand(verifyCmd)

}
