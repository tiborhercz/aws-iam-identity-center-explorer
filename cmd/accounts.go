package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/internal/export"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/internal/model"
	ssoexplore "github.com/tiborhercz/aws-iam-identity-center-explorer/internal/sso-explore"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/internal/util"
	"log"
)

var (
	accountsOptions model.Options

	accountsCmd = &cobra.Command{
		Use:  "accounts",
		Args: validateOptions,
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFlags(0)

			log.Println("Getting the SSO Accounts information. This may take a while...")
			data := ssoexplore.Explore(accountsOptions.InstanceArn, accountsOptions.IdentityStoreId)

			jsondata, _ := json.Marshal(data)

			if accountsOptions.ExportType == "json" {
				export.JsonFile(jsondata)
			} else if accountsOptions.ExportType == "csv" {
				//export.CsvFile(jsondata)
			} else {
				prettyJson, _ := util.PrettifyJson(jsondata)
				log.Printf("\n%v", prettyJson)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(accountsCmd)
	accountsCmd.Flags().StringVar(&accountsOptions.InstanceArn, "instanceArn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	accountsCmd.Flags().StringVar(&accountsOptions.IdentityStoreId, "identityStoreId", "", "The globally unique identifier for the identity store.")
	accountsCmd.Flags().StringVar(&accountsOptions.ExportType, "exportType", "", "Export results to a file. Allowed values: 'json' or 'csv'")
}
