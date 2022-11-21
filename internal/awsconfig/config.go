package awsconfig

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

func InitConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("eu-west-1"),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}
