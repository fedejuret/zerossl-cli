/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/api/structs/requests"
	"github.com/fedejuret/zerossl-golang-cli/lib/csr"
	"github.com/fedejuret/zerossl-golang-cli/lib/models"
	certificate_service "github.com/fedejuret/zerossl-golang-cli/lib/services"
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

		spinner := utils.GetSpinner("Creating certificate for "+commonName, "fgGreen")

		spinner.Start()
		response := api.Post("/certificates", createCertificateRequest)
		spinner.Stop()

		certificate, err := models.UnmarshalCertificate(response)

		if err != nil {
			log.Fatal(err)
		}

		Box := box.New(box.Config{Px: 10, Py: 2, Type: "Classic", Color: "Green", TitlePos: "Inside"})
		Box.Print("Right!", "Certificate for "+commonName+" has been created with ID: "+certificate.ID)

		validateMethod := -1

		validateMethod, _, _ = utils.GetSelectPromt("How do you want to validate your domain?", []string{"Email verification", "File upload", "Add CNAME record to DNS"})

		if validateMethod == 1 { // File upload

			uploadFileUrl, err := certificate.GetFileValidationURLHTTPS()

			if err != nil {
				log.Fatal(err)
			}

			parsedURL, err := url.Parse(uploadFileUrl)
			if err != nil {
				log.Fatal(err)
			}

			fileName := path.Base(parsedURL.Path)
			fileContent, err := certificate.GetFileValidationContent()

			if err != nil {
				log.Fatal(err)
			}

			fileData := strings.Join(fileContent, "")

			err = os.WriteFile(fileName, []byte(fileData), 0664)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(color.YellowString("Almost done!"))
			fmt.Println(color.CyanString("The file " + fileName + " was created that you must upload to the following path: " + uploadFileUrl))
		} else if validateMethod == 2 {

			cname, content, err := certificate.GetDNSValidation()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(color.YellowString("Almost done!"))
			fmt.Println(color.CyanString("To continue you have to create the following records in your DNS"))

			fmt.Println("")

			fmt.Println(color.CyanString("Type: "), "CNAME")
			fmt.Println(color.CyanString("Name: "), cname)
			fmt.Println(color.CyanString("Content: "), content)
		}

		certificate_service.Store(certificate, int8(validateMethod))

	},
}

func init() {
	certificatesCmd.AddCommand(createCmd)
}
