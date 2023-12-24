/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/api/structs/requests"
	"github.com/fedejuret/zerossl-golang-cli/lib/csr"
	"github.com/fedejuret/zerossl-golang-cli/lib/utils"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new certificate",
	Run: func(cmd *cobra.Command, args []string) {

		commonName, _ := utils.GetStringPromt("Which domain do you want to create a certificate for?")

		csrPromt := promptui.Select{
			Label:  color.CyanString("Do you want to complete the CSR data by yourself?"),
			Stdout: &utils.BellSkipper{},
			Items: []string{
				"Yes",
				"No, complete automatically",
			},
		}

		csrPromtResponse, _, err := csrPromt.Run()

		if err != nil {
			return
		}

		var csrGenerateStruct csr.Generate

		if csrPromtResponse == 0 {
			// Ask for other CSR dat
			organization, _ := utils.GetStringPromt("Organization: ")
			organizationUnit, _ := utils.GetStringPromt("Organization unit: ")
			country, _ := utils.GetStringPromt("Country in two digits. Example: [AR]: ")
			state, _ := utils.GetStringPromt("State: ")
			locality, _ := utils.GetStringPromt("Locality: ")

			csrGenerateStruct = csr.Generate{
				CommonName:       commonName,
				Organization:     organization,
				Country:          country,
				State:            state,
				Locality:         locality,
				OrganizationUnit: organizationUnit,
			}

		} else {
			csrGenerateStruct = csr.Generate{
				CommonName:       commonName,
				Organization:     "FedejuretONG",
				Country:          "AR",
				State:            "La Pampa",
				Locality:         "Guatraché",
				OrganizationUnit: "Development",
			}

		}

		csrBytes, err := csrGenerateStruct.Create()

		certificateLongTime, _, _ := utils.GetSelectPromt("How long do you want to create the certificate for?", []string{"90 days", "365 days"})
		var certificateTime uint16
		if certificateLongTime == 0 {
			certificateTime = 90
		} else {
			certificateTime = 365
		}

		createCertificateRequest := &requests.CreateCertificationStructure{
			Domains:      commonName,
			Csr:          string(csrBytes),
			ValidityDays: certificateTime,
		}

		response := api.Post("/certificates", createCertificateRequest)
		fmt.Println(string(response))

	},
}

func init() {
	certificatesCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
