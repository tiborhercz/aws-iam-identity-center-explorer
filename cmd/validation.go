package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func validateOptions(cmd *cobra.Command, args []string) error {
	if rootOptions.Region == "" {
		log.Fatalln("--region should be provided")
	}

	if rootOptions.IdentityStoreId == "" {
		log.Fatalln("--identityStoreId should be provided")
	}

	if rootOptions.InstanceArn == "" {
		log.Fatalln("--instanceArn should be provided")
	}

	if groupsOptions.ExportType != "" && !(groupsOptions.ExportType == "json" || groupsOptions.ExportType == "csv") {
		log.Fatalln("--exportType must either be 'json' or 'csv'")
	}

	return nil
}
