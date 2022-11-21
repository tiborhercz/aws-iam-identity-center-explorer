package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/internal/model"
	"os"
)

var (
	rootOptions model.RootOptions

	rootCmd = &cobra.Command{
		Use: "aws-iam-identity-center-explorer",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootOptions.Region, "region", "", "AWS region where the IAM Identity Center instance is located.")
	rootCmd.PersistentFlags().StringVar(&rootOptions.InstanceArn, "instanceArn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	rootCmd.PersistentFlags().StringVar(&rootOptions.IdentityStoreId, "identityStoreId", "", "The globally unique identifier for the identity store.")
}
