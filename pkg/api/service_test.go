package api_test

import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/golang/mock/gomock"
	api "github.com/gregfurman/docker-ecr/pkg/api"
	docker "github.com/gregfurman/docker-ecr/pkg/docker/mock_docker"
	ecr "github.com/gregfurman/docker-ecr/pkg/ecr/mock_ecr"
)

var errExpected = errors.New("error")

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dockerSvc := docker.NewMockService(ctrl)
	ecrSvc := ecr.NewMockService(ctrl)
	api := api.NewService(dockerSvc, ecrSvc)

	success := func(t *testing.T) {
		auth := base64.StdEncoding.EncodeToString([]byte("user:pass"))
		ecrSvc.EXPECT().GetAuth().Return(&types.AuthorizationData{AuthorizationToken: &auth}, nil)

		fmtAuth, err := api.Login()
		if err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}

		expectedAuth := base64.StdEncoding.EncodeToString([]byte("{\"Username\":\"user\",\"Password\":\"pass\"}"))
		if expectedAuth != *fmtAuth {
			t.Errorf("Authentication string was incorrect. Expected %s, got %s", expectedAuth, *fmtAuth)
		}
	}

	failedAuthFmt := func(t *testing.T) {
		auth := base64.StdEncoding.EncodeToString([]byte("user"))
		ecrSvc.EXPECT().GetAuth().Return(&types.AuthorizationData{AuthorizationToken: &auth}, nil)

		fmtAuth, err := api.Login()

		if err == nil {
			t.Error("Expected an error, got nil")
		}

		if fmtAuth != nil {
			t.Errorf("Expected nil, received %s", *fmtAuth)
		}
	}

	failedClientCall := func(t *testing.T) {
		ecrSvc.EXPECT().GetAuth().Return(nil, errExpected)

		fmtAuth, err := api.Login()

		if err == nil {
			t.Error("Expected an error, got nil")
		}

		if fmtAuth != nil {
			t.Errorf("Expected nil, received %s", *fmtAuth)
		}
	}

	t.Run("successful auth", success)
	t.Run("failed client call", failedClientCall)
	t.Run("failed auth formatting", failedAuthFmt)
}

func TestPull(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dockerSvc := docker.NewMockService(ctrl)
	ecrSvc := ecr.NewMockService(ctrl)
	api := api.NewService(dockerSvc, ecrSvc)

	ref := "012345678910.dkr.ecr.xx-region-1.amazonaws.com/test/repo"

	auth := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	expectedAuth := base64.StdEncoding.EncodeToString([]byte("{\"Username\":\"user\",\"Password\":\"pass\"}"))

	successfulPull := func(t *testing.T) {
		ecrSvc.EXPECT().GetAuth().Return(&types.AuthorizationData{AuthorizationToken: &auth}, nil)
		dockerSvc.EXPECT().Pull(ref, expectedAuth).Return(nil)

		err := api.Pull(ref)
		if err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}
	}

	failedPullWithError := func(t *testing.T) {
		ecrSvc.EXPECT().GetAuth().Return(&types.AuthorizationData{AuthorizationToken: &auth}, nil)
		dockerSvc.EXPECT().Pull(ref, expectedAuth).Return(errExpected)

		err := api.Pull(ref)
		if !errors.Is(err, errExpected) {
			t.Errorf("Unexpected error occurred. Expected %v, got %v", errExpected, err)
		}
	}

	t.Run("successful pull", successfulPull)
	t.Run("failed pull", failedPullWithError)
}

func TestPush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dockerSvc := docker.NewMockService(ctrl)
	ecrSvc := ecr.NewMockService(ctrl)
	api := api.NewService(dockerSvc, ecrSvc)

	repoName := "test/repo"
	isImageTagsMutable := true
	repoTags := map[string]string{"test": "test"}
	ref := "012345678910.dkr.ecr.xx-region-1.amazonaws.com/test/repo"

	imageTags := []string{"hash_123", "test"}
	repo := types.Repository{RepositoryUri: &ref}

	auth := base64.StdEncoding.EncodeToString([]byte("user:pass"))
	expectedAuth := base64.StdEncoding.EncodeToString([]byte("{\"Username\":\"user\",\"Password\":\"pass\"}"))

	successfulPush := func(t *testing.T) {
		ecrSvc.EXPECT().CreateEcrRepository(repoName, isImageTagsMutable, repoTags).Return(&repo, nil)
		ecrSvc.EXPECT().GetAuth().Return(&types.AuthorizationData{AuthorizationToken: &auth}, nil)

		gomock.InOrder(
			dockerSvc.EXPECT().Push(*repo.RepositoryUri+":hash_123", expectedAuth).Return(nil),
			dockerSvc.EXPECT().Push(*repo.RepositoryUri+":test", expectedAuth).Return(nil),
		)

		err := api.Push(repoName, repoTags, imageTags...)
		if err != nil {
			t.Errorf("Unexpected error occurred. Expected nil, got %v", err)
		}
	}

	t.Run("successful push to docker", successfulPush)
}