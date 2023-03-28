package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type Service interface {
	Pull(imageRefUrl, registryAuth string) error
	Push(imageRefUrl, registryAuth string) error
	Build(dockerfile string, tags ...string) error
	Tag(src, dest string) error
}

type ServiceImpl struct {
	client *client.Client
}

func NewService(client *client.Client) Service {
	return &ServiceImpl{client: client}
}

// Pull requests the docker host to pull an image from a remote repository.
// The full remote image path is required as well as authentication for the registry.
func (s *ServiceImpl) Pull(imageRefUrl, registryAuth string) error {
	if !IsBase64(registryAuth) {
		return fmt.Errorf("error: registry authorisation string in form is not base64 encoded")
	}

	reader, err := s.client.ImagePull(context.TODO(), imageRefUrl, types.ImagePullOptions{RegistryAuth: registryAuth})
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := parse(reader); err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	}

	return nil
}

// Push requests the docker host to push an image to a remote repository.
// The full remote image path is required as well as authentication for the registry.
func (s *ServiceImpl) Push(imageRefUrl, registryAuth string) error {
	if !IsBase64(registryAuth) {
		return fmt.Errorf("error: registry authorisation string in form is not base64 encoded")
	}

	reader, err := s.client.ImagePush(context.TODO(), imageRefUrl, types.ImagePushOptions{RegistryAuth: registryAuth})
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := parse(reader); err != nil {
		return fmt.Errorf("failed to push image: %w", err)
	}

	return nil
}

func (s *ServiceImpl) Build(dockerfile string, tags ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	tar, err := archive.TarWithOptions(dockerfile, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       tags,
		Remove:     true,
	}

	res, err := s.client.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := parse(res.Body); err != nil {
		return fmt.Errorf("Failed to build image: %w", err)
	}

	return nil
}

func (s *ServiceImpl) Tag(src, dest string) error {
	return s.client.ImageTag(context.TODO(), src, dest)
}
