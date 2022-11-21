package ssoadmin

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"log"
)

func ListPermissionSetsProvisionedToAccount(awscfg aws.Config, instanceArn string, account orgTypes.Account) []string {
	ssoadminClient := ssoadmin.NewFromConfig(awscfg)

	var permissionSets []string
	var nextToken *string
	isFinished := false
	for !isFinished {
		ress, err := ssoadminClient.ListPermissionSetsProvisionedToAccount(context.TODO(), &ssoadmin.ListPermissionSetsProvisionedToAccountInput{
			AccountId:   account.Id,
			InstanceArn: &instanceArn,
		})

		if err != nil {
			log.Fatal(err)
		}

		nextToken = ress.NextToken

		for _, permissionSet := range ress.PermissionSets {
			permissionSets = append(permissionSets, permissionSet)
		}

		if nextToken == nil {
			isFinished = true
		}
	}

	return permissionSets
}

func ListAccountAssignments(awscfg aws.Config, account orgTypes.Account, instanceArn string, permissionSetArn string) []string {
	ssoadminClient := ssoadmin.NewFromConfig(awscfg)

	var accountAssignments []string
	var nextToken *string
	isFinished := false
	for !isFinished {
		ress, err := ssoadminClient.ListAccountAssignments(context.TODO(), &ssoadmin.ListAccountAssignmentsInput{
			AccountId:        account.Id,
			InstanceArn:      &instanceArn,
			PermissionSetArn: &permissionSetArn,
		})

		if err != nil {
			log.Fatal(err)
		}

		nextToken = ress.NextToken

		for _, accountAssignment := range ress.AccountAssignments {
			if accountAssignment.PrincipalType == "GROUP" {
				accountAssignments = append(accountAssignments, *accountAssignment.PrincipalId)
			}
		}

		if nextToken == nil {
			isFinished = true
		}
	}

	return accountAssignments
}
