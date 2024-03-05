package docker_test

import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/golang/mock/gomock"
	docker "github.com/gregfurman/docker-ecr/pkg/docker"
	mock "github.com/gregfurman/docker-ecr/pkg/docker/mock_docker"
)

//nolint:gochecknoglobals
var (
	auth        = base64.StdEncoding.EncodeToString([]byte("user:pass"))
	ref         = "docker-registry.com/test/repo"
	errExpected = errors.New("error")
)

func TestPull(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	svc := docker.NewService(client)
	options := types.ImagePullOptions{RegistryAuth: auth}

	success := func(t *testing.T) {
		client.EXPECT().Pull(ref, options).Return(nil)

		if err := svc.Pull(ref, auth); err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}
	}

	failedPullWithError := func(t *testing.T) {
		client.EXPECT().Pull(ref, options).Return(errExpected)

		if err := svc.Pull(ref, auth); err == nil {
			t.Error("Expected error, got nil")
		}
	}

	t.Run("successful pull", success)
	t.Run("failed pull with error", failedPullWithError)
}

func TestPush(t *testing.T) {
	ref := "docker-registry.com/test/repo"
	auth := base64.StdEncoding.EncodeToString([]byte("user:pass"))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	svc := docker.NewService(client)
	options := types.ImagePushOptions{RegistryAuth: auth}

	success := func(t *testing.T) {
		client.EXPECT().Push(ref, options).Return(nil)

		if err := svc.Push(ref, auth); err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}
	}

	failedPushWithError := func(t *testing.T) {
		client.EXPECT().Push(ref, options).Return(errExpected)

		if err := svc.Push(ref, auth); err == nil {
			t.Error("Expected error, got nil")
		}
	}

	t.Run("successful push", success)
	t.Run("failed push with error", failedPushWithError)
}

func TestBuild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	svc := docker.NewService(client)

	dockerfile := "dockerfile"
	tags := []string{"test", "image", "tags"}
	options := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       tags,
		Remove:     true,
	}

	success := func(t *testing.T) {
		client.EXPECT().Build(options).Return(nil)

		if err := svc.Build(dockerfile, tags...); err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}
	}

	failedBuildWithError := func(t *testing.T) {
		client.EXPECT().Build(options).Return(errExpected)

		if err := svc.Build(dockerfile, tags...); err == nil {
			t.Error("Expected error, got nil")
		}
	}

	t.Run("successful build", success)
	t.Run("failed build with error", failedBuildWithError)
}

func TestTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	svc := docker.NewService(client)

	src := "imageName"
	dest := "tag"

	success := func(t *testing.T) {
		client.EXPECT().Tag(src, dest).Return(nil)

		if err := svc.Tag(src, dest); err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}
	}

	failedTagWithError := func(t *testing.T) {
		client.EXPECT().Tag(src, dest).Return(errExpected)

		if err := svc.Tag(src, dest); err == nil {
			t.Error("Expected error, got nil")
		}
	}

	t.Run("successful tag", success)
	t.Run("failed tag with error", failedTagWithError)
}
