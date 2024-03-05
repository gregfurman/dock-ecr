package docker

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func NewClient(opts ...client.Opt) Client {
	dockerClient, err := client.NewClientWithOpts(opts...)
	if err != nil {
		logrus.Fatalf("Failed to load new Docker client: %s", err)
		return nil
	}

	dockerClient.NegotiateAPIVersion(context.TODO())

	return &ClientImpl{client: dockerClient}
}
