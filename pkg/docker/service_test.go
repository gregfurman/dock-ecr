package docker

import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/golang/mock/gomock"
	docker "github.com/gregfurman/docker-ecr/pkg/docker/mock_docker"
)

func TestDocker(t *testing.T) {
	setupMock := func() *docker.MockClient {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		return docker.NewMockClient(ctrl)
	}

	ref := "docker-registry.com/test/repo"
	auth := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	options := types.ImagePullOptions{RegistryAuth: auth}

	TestPull := func(t *testing.T) {
		client := setupMock()
		svc := NewService(client)

		success := func(t *testing.T) {
			client.EXPECT().Pull(ref, options).Return(nil)

			if err := svc.Pull(ref, auth); err != nil {
				t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
			}
		}

		failedPullWithError := func(t *testing.T) {
			client.EXPECT().Pull(ref, options).Return(errors.New("error"))

			if err := svc.Pull(ref, auth); err == nil {
				t.Error("Expected error, got nil")
			}
		}

		t.Run("successful pull", success)
		t.Run("failed pull with error", failedPullWithError)
	}

	TestPush := func(t *testing.T) {
		client := setupMock()
		svc := NewService(client)

		success := func(t *testing.T) {
			client.EXPECT().Push(ref, options).Return(nil)

			if err := svc.Push(ref, auth); err != nil {
				t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
			}
		}

		failedPushWithError := func(t *testing.T) {
			client.EXPECT().Push(ref, options).Return(errors.New("error"))

			if err := svc.Push(ref, auth); err == nil {
				t.Error("Expected error, got nil")
			}
		}

		t.Run("successful push", success)
		t.Run("failed push with error", failedPushWithError)
	}

	TestBuild := func(t *testing.T) {
		client := setupMock()
		svc := NewService(client)

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
			client.EXPECT().Build(options).Return(errors.New("error"))

			if err := svc.Build(dockerfile, tags...); err == nil {
				t.Error("Expected error, got nil")
			}
		}

		t.Run("successful build", success)
		t.Run("failed build with error", failedBuildWithError)
	}

	TestTag := func(t *testing.T) {
		client := setupMock()
		svc := NewService(client)

		src := "imageName"
		dest := "tag"

		success := func(t *testing.T) {
			client.EXPECT().Tag(src, dest).Return(nil)

			if err := svc.Tag(src, dest); err != nil {
				t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
			}
		}

		failedTagWithError := func(t *testing.T) {
			client.EXPECT().Tag(src, dest).Return(errors.New("error"))

			if err := svc.Tag(src, dest); err == nil {
				t.Error("Expected error, got nil")
			}
		}

		t.Run("successful tag", success)
		t.Run("failed tag with error", failedTagWithError)
	}

	t.Run("TestBuild", TestBuild)
	t.Run("TestPull", TestPull)
	t.Run("TestPush", TestPush)
	t.Run("TestTag", TestTag)
}
