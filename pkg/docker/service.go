package docker

import (
	"errors"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
)

type Service interface {
	Pull(imageRefURL, registryAuth string) error
	Push(imageRefURL, registryAuth string) error
	Build(dockerfile string, tags ...string) error
	Tag(src, dest string) error
	Login(auth string) error
}

type ServiceImpl struct {
	client Client
}

func NewService(client Client) *ServiceImpl {
	return &ServiceImpl{client: client}
}

var errNotBase64 = errors.New("registry authorisation string in form is not base64 encoded")

// Pull requests the docker host to pull an image from a remote repository.
// The full remote image path is required as well as authentication for the registry.
func (s *ServiceImpl) Pull(imageRefURL, registryAuth string) error {
	if !IsBase64(registryAuth) {
		return errNotBase64
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
		return errNotBase64
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
	return s.client.Tag(src, dest)
}

func (s *ServiceImpl) Login(auth string) error {
	return s.client.Login(registry.AuthConfig{Auth: auth})
}
