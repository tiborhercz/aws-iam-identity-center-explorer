package awsconfig

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

func InitConfig(region string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}
