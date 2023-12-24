/*
Copyright © 2024 Federico Juretich <fedejuret@gmail.com>
*/
package certificates

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fedejuret/zerossl-golang-cli/lib/api"
	"github.com/fedejuret/zerossl-golang-cli/lib/api/structs/responses"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your certificates",
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")
		cname, _ := cmd.Flags().GetString("cname")

		queryParams := make(map[string]string)
		queryParams["certificate_status"] = status
		queryParams["search"] = cname

		response := api.Get("/certificates", queryParams)

		certificates, err := responses.UnmarshalCertificates(response)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		text := fmt.Sprintf("Certificates found: %d\n", certificates.TotalCount)
		fmt.Println(text)

		if certificates.TotalCount == 0 {
			return
		}

		headerFmt := color.New(color.FgGreen, color.Underline, color.Bold).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("ID", "Name", "Status", "Expiration")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, cert := range certificates.Results {
			tbl.AddRow(cert.ID, cert.CommonName, cert.Status, cert.Expires)
		}

		tbl.Print()
	},
}

func init() {
	certificatesCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("status", "s", "issued", "Certificate status. Possible values: draft, pending_validation, issued, cancelled, revoked, expired")
	listCmd.Flags().StringP("cname", "c", "", "Use this parameter to search for certificates having the given common name or SAN")
}
