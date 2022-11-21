package organization

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"log"
)

func ListAllAccounts(awscfg aws.Config) []types.Account {
	org := organizations.NewFromConfig(awscfg)

	var accountsList []types.Account
	var nextToken *string
	isFinished := false
	for !isFinished {
		accounts, err := org.ListAccounts(context.TODO(), &organizations.ListAccountsInput{
			NextToken: nextToken,
		})

		if err != nil {
			log.Fatal(err)
		}

		nextToken = accounts.NextToken

		for _, account := range accounts.Accounts {
			accountsList = append(accountsList, account)
		}

		if nextToken == nil {
			isFinished = true
		}
	}

	return accountsList
}
