/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-cli/lib/api"
	"github.com/fedejuret/zerossl-cli/lib/api/structs/requests"
	"github.com/fedejuret/zerossl-cli/lib/models"
	"github.com/fedejuret/zerossl-cli/lib/utils"
	"github.com/spf13/cobra"
)

// revokeCmd represents the revoke command
var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke certificate",
	Long:  `Revoking a certificate is a command that is only available for issued certificates`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		certId := args[0]

		spinner := utils.GetSpinner("Fetching certificate...", "fgYellow")

		spinner.Start()
		response := api.Get("/certificates/"+certId, map[string]string{})
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
						fmt.Println(color.RedString("Error when trying to get certificate: " + errorMessage))
					} else {
						fmt.Println(color.RedString("Error when trying to get certificate: unknown error type"))
					}
				} else {
					fmt.Println(color.RedString("Error when trying to get certificate: unknown error information"))
				}

			}
		} else {

			certificate, err := models.UnmarshalCertificate(response)

			if certificate.Status != "issued" {
				fmt.Println(color.RedString("Only issued certificates can be rovoked"))
				return
			}

			if err != nil {
				log.Fatal(err)
			}

			revokeReason, _, err := utils.GetSelectPromt("Enter your reason", []string{"Compromised private key", "Subjects' name or identity information has changed", "Certificate has been replaced", "Authorized domain names are no longer owned", "Other"})

			realReason := "Unspecified"

			switch revokeReason {
			case 0:
				realReason = "keyCompromise"
				break
			case 1:
				realReason = "affiliationChanged"
				break
			case 2:
				realReason = "Superseded"
				break
			case 3:
				realReason = "cessationOfOperation"
				break
			case 4:
				realReason = "Unspecified"
				break
			}

			spinner := utils.GetSpinner("Executing...", "fgYellow")

			spinner.Start()
			response = api.Post("/certificates/"+certificate.ID+"/revoke", requests.RevokeCertificateStructure{
				Reason: realReason,
			})
			spinner.Stop()

			Box := box.New(box.Config{Px: 10, Py: 2, Type: "Classic", Color: "Green", TitlePos: "Inside"})
			Box.Print("Right!", "Certificate revoked successfully")
		}
	},
}

func init() {
	certificatesCmd.AddCommand(revokeCmd)
}
