package docker

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type Client interface {
	Build(options types.ImageBuildOptions) error
	Pull(refStr string, options types.ImagePullOptions) error
	Push(image string, options types.ImagePushOptions) error
	Tag(source string, target string) error
	Login(auth registry.AuthConfig) error
}

type ClientImpl struct {
	client *client.Client
}

func (c *ClientImpl) Build(options types.ImageBuildOptions) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory path: %w", err)
	}

	tar, err := archive.TarWithOptions(wd, &archive.TarOptions{})
	if err != nil {
		return fmt.Errorf("failed to create an archive of Dockerfile %s: %w", options.Dockerfile, err)
	}

	res, err := c.client.ImageBuild(context.Background(), tar, options)
	if err != nil {
		return fmt.Errorf("build image request to Docker daemon failed: %w", err)
	}
	defer res.Body.Close()

	return parseDockerOutput(res.Body)
}

func (c *ClientImpl) Pull(refStr string, options types.ImagePullOptions) error {
	reader, err := c.client.ImagePull(context.Background(), refStr, options)
	if err != nil {
		return fmt.Errorf("failed to pull image %s: %w", refStr, err)
	}
	defer reader.Close()

	return parseDockerOutput(reader)
}

func (c *ClientImpl) Push(image string, options types.ImagePushOptions) error {
	reader, err := c.client.ImagePush(context.TODO(), image, options)
	if err != nil {
		return fmt.Errorf("failed to push image %s to remote registry: %w", image, err)
	}
	defer reader.Close()

	return parseDockerOutput(reader)
}

func (c *ClientImpl) Tag(source string, target string) error {
	err := c.client.ImageTag(context.Background(), source, target)
	if err != nil {
		return fmt.Errorf("failed to tag image %s as %s: %w", source, target, err)
	}

	return nil
}

func (c *ClientImpl) Login(auth registry.AuthConfig) error {
	_, err := c.client.RegistryLogin(context.TODO(), auth)
	if err != nil {
		return fmt.Errorf("login request to registry failed: %w", err)
	}

	return nil
}
