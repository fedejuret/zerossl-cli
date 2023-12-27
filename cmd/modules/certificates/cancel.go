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
	"github.com/fedejuret/zerossl-cli/lib/utils"
	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel any certificate",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		certId := args[0]

		spinner := utils.GetSpinner("Executing...", "fgYellow")

		spinner.Start()
		response := api.Post("/certificates/"+certId+"/cancel", []string{})
		spinner.Stop()

		var responseJson map[string]interface{}

		if err := json.Unmarshal(response, &responseJson); err != nil {
			log.Fatal(err)
		}

		success, successExists := responseJson["success"]
		if successExists {

			if reflect.TypeOf(success).Kind() == reflect.Float64 {
				Box := box.New(box.Config{Px: 10, Py: 2, Type: "Classic", Color: "Green", TitlePos: "Inside"})
				Box.Print("Right!", "Certificate cancelled successfully")
			} else {
				errorInfo, errorExists := responseJson["error"].(map[string]interface{})
				if errorExists {
					errorMessage, typeExists := errorInfo["type"].(string)
					if typeExists {
						fmt.Println(color.RedString("Error when trying to cancel certificate: " + errorMessage))
					} else {
						fmt.Println(color.RedString("Error when trying to cancel certificate: unknown error type"))
					}
				} else {
					fmt.Println(color.RedString("Error when trying to cancel certificate: unknown error information"))
				}
			}
		} else {
			fmt.Println(color.RedString("Error: 'success' key not found in response"))
		}
	},
}

func init() {
	certificatesCmd.AddCommand(cancelCmd)
}
