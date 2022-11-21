package sso_explore

import (
	"fmt"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/internal/awsconfig"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/pkg/aws/identitystore"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/pkg/aws/organization"
	"github.com/tiborhercz/aws-iam-identity-center-explorer/pkg/aws/ssoadmin"
)

type SsoAccountInformation struct {
	AccountName string
	Groups      []string
}

type SsoGroupInformation struct {
	AccountName string
	AccountId   string
}

func Explore(instanceArn string, identityStoreId string, region string) map[string]SsoAccountInformation {
	awscfg := awsconfig.InitConfig(region)

	groupsAttachedToAccounts := make(map[string]SsoAccountInformation)

	allAccounts := organization.ListAllAccounts(awscfg)
	fmt.Printf("Found %v accounts \n", len(allAccounts))

	iter := 1
	for _, account := range allAccounts {
		fmt.Printf("\rOn %v/%d accounts", iter, len(allAccounts))
		ssoAccountInformation := &SsoAccountInformation{}

		ssoAccountInformation.AccountName = *account.Name

		for _, permissionSetArn := range ssoadmin.ListPermissionSetsProvisionedToAccount(awscfg, instanceArn, account) {

			for _, GroupId := range ssoadmin.ListAccountAssignments(awscfg, account, instanceArn, permissionSetArn) {
				group, err := identitystore.DescribeGroup(awscfg, identityStoreId, GroupId)
				if err != nil {
					continue
				}

				ssoAccountInformation.Groups = append(ssoAccountInformation.Groups, *group.DisplayName)
			}
		}

		groupsAttachedToAccounts[*account.Id] = *ssoAccountInformation

		iter++
	}

	return groupsAttachedToAccounts
}

func TransformDataToGroups(data map[string]SsoAccountInformation) map[string][]SsoGroupInformation {
	accountsAttachedToGroups := make(map[string][]SsoGroupInformation)

	for index, accountInformation := range data {
		for _, groupName := range accountInformation.Groups {
			accountsAttachedToGroups[groupName] = append(accountsAttachedToGroups[groupName], SsoGroupInformation{
				AccountName: accountInformation.AccountName,
				AccountId:   index,
			})
		}
	}

	return accountsAttachedToGroups
}
