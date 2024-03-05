package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
)

type Service interface {
	Pull(imageRefURL, registryAuth string) error
	Push(imageRefURL, registryAuth string) error
	Build(dockerfile string, tags ...string) error
	Tag(src, dest string) error
}

type ServiceImpl struct {
	client Client
}

func NewService(client Client) Service {
	return &ServiceImpl{client: client}
}

// Pull requests the docker host to pull an image from a remote repository.
// The full remote image path is required as well as authentication for the registry.
func (s *ServiceImpl) Pull(imageRefURL, registryAuth string) error {
	if !IsBase64(registryAuth) {
		return fmt.Errorf("error: registry authorisation string in form is not base64 encoded")
	}

	if err := s.client.Pull(imageRefURL, types.ImagePullOptions{RegistryAuth: registryAuth}); err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	}

	return nil
}

// Push requests the docker host to push an image to a remote repository.
// The full remote image path is required as well as authentication for the registry.
func (s *ServiceImpl) Push(imageRefURL, registryAuth string) error {
	if !IsBase64(registryAuth) {
		return fmt.Errorf("error: registry authorisation string in form is not base64 encoded")
	}

	if err := s.client.Push(imageRefURL, types.ImagePushOptions{RegistryAuth: registryAuth}); err != nil {
		return fmt.Errorf("failed to push image: %w", err)
	}

	return nil
}

func (s *ServiceImpl) Build(dockerfile string, tags ...string) error {
	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       tags,
		Remove:     true,
	}

	if err := s.client.Build(opts); err != nil {
		return fmt.Errorf("failed to build image with build options %v: %w", opts, err)
	}

	return nil
}

func (s *ServiceImpl) Tag(src, dest string) error {
	if err := s.client.Tag(src, dest); err != nil {
		return fmt.Errorf("failed to tag image '%s' with tag '%s': %w", src, dest, err)
	}

	return nil
}
