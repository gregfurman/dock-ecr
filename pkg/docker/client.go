package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type Client interface {
	Build(options types.ImageBuildOptions) error
	Pull(refStr string, options types.ImagePullOptions) error
	Push(image string, options types.ImagePushOptions) error
	Tag(source string, target string) error
}

type ClientImpl struct {
	client *client.Client
}

func (c *ClientImpl) Build(options types.ImageBuildOptions) error {
	tar, err := archive.TarWithOptions(options.Dockerfile, &archive.TarOptions{})
	if err != nil {
		return err
	}

	res, err := c.client.ImageBuild(context.Background(), tar, options)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return parse(res.Body)
}

func (c *ClientImpl) Pull(refStr string, options types.ImagePullOptions) error {
	reader, err := c.client.ImagePull(context.Background(), refStr, options)
	if err != nil {
		return err
	}
	defer reader.Close()

	return parse(reader)
}

func (c *ClientImpl) Push(image string, options types.ImagePushOptions) error {
	reader, err := c.client.ImagePush(context.TODO(), image, options)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := parse(reader); err != nil {
		return err
	}

	return nil
}

func (c *ClientImpl) Tag(source string, target string) error {
	return c.client.ImageTag(context.Background(), source, target)
}
