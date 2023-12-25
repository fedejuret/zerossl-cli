/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/api/structs/responses"
	"github.com/fedejuret/zerossl-golang-cli/lib/utils"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your certificates",
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")
		cname, _ := cmd.Flags().GetString("cname")
		expiringDays, _ := cmd.Flags().GetInt("expiring-days")

		queryParams := make(map[string]string)
		queryParams["certificate_status"] = status

		if len(cname) > 0 {
			queryParams["search"] = cname
		}

		spinner := utils.GetSpinner("Fetching certificates", "fgYellow")

		spinner.Start()
		response := api.Get("/certificates", queryParams)
		spinner.Stop()

		certificates, err := responses.UnmarshalCertificates(response)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		headerFmt := color.New(color.FgGreen, color.Underline, color.Bold).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("ID", "Name", "Status", "Expiration")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		totalCerts := 0

		for _, cert := range certificates.Results {
			mustPrint := true
			certExpiration, err := time.Parse("2006-01-02 15:04:05", cert.Expires)
			if expiringDays != -1 {

				if err != nil {
					log.Fatal(err)
				}

				expirationDate := time.Now().Add(time.Duration(expiringDays) * 24 * time.Hour)

				if certExpiration.Before(expirationDate) == false {
					mustPrint = false
				}
			}

			if mustPrint {
				totalCerts++
				tbl.AddRow(cert.ID, cert.CommonName, cert.Status, cert.Expires+" ("+utils.GetTimeAgo().Format(certExpiration)+")")
			}

		}

		text := fmt.Sprintf("Certificates found: %d\n", totalCerts)
		fmt.Println(text)

		if totalCerts == 0 {
			return
		}

		tbl.Print()
	},
}

func init() {
	certificatesCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("status", "s", "issued", "Certificate status. Possible values: draft, pending_validation, expiring_soon, issued, cancelled, revoked, expired")
	listCmd.Flags().StringP("cname", "c", "", "Use this parameter to search for certificates having the given common name or SAN")
	listCmd.Flags().Int("expiring-days", -1, "Get certificates that expire within the specified date")
}
