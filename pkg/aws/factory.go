package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sirupsen/logrus"
)

type loadConfig func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error)

func generateConfig(ctx context.Context, cfgLoader loadConfig, optFns ...func(*config.LoadOptions) error) aws.Config {
	cfg, err := cfgLoader(ctx, optFns...)
	if err != nil {
		logrus.Fatalf("Failed to create config: %s", err)

		return aws.Config{}
	}

	return cfg
}

func NewClient() *ecr.Client {
	return ecr.NewFromConfig(generateConfig(context.TODO(), config.LoadDefaultConfig))
}

func NewClientFromEnv(profile string) *ecr.Client {
	return ecr.NewFromConfig(generateConfig(context.TODO(), config.LoadDefaultConfig, config.WithSharedConfigProfile(profile)))
}

func NewStsClient() *sts.Client {
	return sts.NewFromConfig(generateConfig(context.TODO(), config.LoadDefaultConfig))
}
