/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/models"
	"github.com/fedejuret/zerossl-golang-cli/lib/utils"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
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

			if err != nil {
				log.Fatal(err)
			}

			spinner = utils.GetSpinner("Downloading...", "fgYellow")

			spinner.Start()
			response = api.Get("/certificates/"+certificate.ID+"/download/return", map[string]string{})
			spinner.Stop()

			var responseJson map[string]string

			if err := json.Unmarshal(response, &responseJson); err != nil {
				log.Fatal(err)
			}

			crt := responseJson["certificate.crt"]
			bundle := responseJson["ca_bundle.crt"]

			file, err := os.Create(os.Getenv("ZEROSSL_FOLDER") + "/" + certificate.CommonName + "/certificate.crt")
			if err != nil {
				log.Fatal(err)
			}
			file.WriteString(crt)
			defer file.Close()

			file, err = os.Create(os.Getenv("ZEROSSL_FOLDER") + "/" + certificate.CommonName + "/ca_bundle.crt")
			if err != nil {
				log.Fatal(err)
			}
			file.WriteString(bundle)
			defer file.Close()

			fmt.Println(color.GreenString("Certificate downloaded in"), certificate.CommonName, color.GreenString("folder"))
		}

	},
}

func init() {
	certificatesCmd.AddCommand(downloadCmd)
}
