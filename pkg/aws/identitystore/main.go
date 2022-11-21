package identitystore

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
)

func DescribeGroup(awscfg aws.Config, identityStoreId string, principalId string) (*identitystore.DescribeGroupOutput, error) {
	idsr := identitystore.NewFromConfig(awscfg)

	group, err := idsr.DescribeGroup(context.TODO(), &identitystore.DescribeGroupInput{
		GroupId:         &principalId,
		IdentityStoreId: &identityStoreId,
	})
	if err != nil {
		return nil, err
	}

	return group, nil
}
