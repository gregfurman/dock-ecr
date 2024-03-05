package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/sirupsen/logrus"
)

type loadConfig func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error)

func clientFromConfig(ctx context.Context, cfgLoader loadConfig, optFns ...func(*config.LoadOptions) error) *ecr.Client {
	cfg, err := cfgLoader(ctx, optFns...)
	if err != nil {
		logrus.Fatalf("Failed to create ECR client from config: %s", err)

		return nil
	}

	ecrClient := ecr.NewFromConfig(cfg)

	return ecrClient
}

func NewClient() *ecr.Client {
	return clientFromConfig(context.TODO(), config.LoadDefaultConfig)
}

func NewClientFromEnv(profile string) *ecr.Client {
	return clientFromConfig(context.TODO(), config.LoadDefaultConfig, config.WithSharedConfigProfile(profile))
}
