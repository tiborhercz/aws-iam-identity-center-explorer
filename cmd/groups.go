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
	groupsOptions model.Options

	groupsCmd = &cobra.Command{
		Use:  "groups",
		Args: validateOptions,
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFlags(0)

			log.Println("Getting the SSO Groups information. This may take a while...")

			data := ssoexplore.Explore(groupsOptions.InstanceArn, groupsOptions.IdentityStoreId)

			groupsData := ssoexplore.TransformDataToGroups(data)

			jsondata, _ := json.Marshal(groupsData)

			if groupsOptions.ExportType == "json" {
				export.JsonFile(jsondata)
			} else if groupsOptions.ExportType == "csv" {
				//export.CsvFile(jsondata)
			} else {
				prettyJson, _ := util.PrettifyJson(jsondata)
				log.Printf("\n%v", prettyJson)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(groupsCmd)
	groupsCmd.Flags().StringVar(&groupsOptions.InstanceArn, "instanceArn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	groupsCmd.Flags().StringVar(&groupsOptions.IdentityStoreId, "identityStoreId", "", "The globally unique identifier for the identity store.")
	groupsCmd.Flags().StringVar(&groupsOptions.ExportType, "exportType", "", "Export results to a file. Allowed values: 'json' or 'csv'")
}
