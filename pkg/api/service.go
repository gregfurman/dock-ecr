package api

import (
	"fmt"
	"strings"

	"github.com/gregfurman/docker-ecr/pkg/docker"
	"github.com/gregfurman/docker-ecr/pkg/ecr"
)

type Service interface {
	Build(imageRefURL string, push bool, repositoryName string, repositoryTags map[string]string, imageTags ...string) error
	Push(repositoryName string, repositoryTags map[string]string, imageTags ...string) error
	Pull(imageRefURL string) error
	Login() (*string, error)
}

type ServiceImpl struct {
	dockerService docker.Service
	ecrService    ecr.Service
}

func NewService(dockerSvc docker.Service, ecrSvc ecr.Service) Service {
	service := ServiceImpl{
		dockerService: dockerSvc,
		ecrService:    ecrSvc,
	}

	return &service
}

func (s *ServiceImpl) Login() (*string, error) {
	auth, err := s.ecrService.GetAuth()
	if err != nil {
		return nil, err
	}

	fmtAuth, err := ecr.FormatAuthDetails(*auth.AuthorizationToken)
	if err != nil {
		return nil, err
	}

	return fmtAuth, nil
}

func (s *ServiceImpl) Build(imageRefURL string, push bool, repositoryName string, repositoryTags map[string]string, imageTags ...string) error {
	repo, err := s.ecrService.CreateEcrRepository(repositoryName, true, repositoryTags)
	if err != nil {
		return err
	}

	for i, tag := range imageTags {
		imageTags[i] = fmt.Sprintf("%s:%s", *repo.RepositoryUri, tag)
	}

	if err := s.dockerService.Build(imageRefURL, imageTags...); err != nil {
		return err
	}

	if !push {
		return nil
	}

	return s.Push(repositoryName, repositoryTags, imageTags...)
}

func (s *ServiceImpl) Push(repositoryName string, repositoryTags map[string]string, imageTags ...string) error {
	auth, err := s.Login()
	if err != nil {
		return err
	}

	repo, err := s.ecrService.CreateEcrRepository(repositoryName, true, repositoryTags)
	if err != nil {
		return err
	}

	for _, tag := range imageTags {
		if uri := *repo.RepositoryUri; !strings.HasPrefix(tag, uri) {
			tag = fmt.Sprintf("%s:%s", uri, tag)
		}
		if err := s.dockerService.Push(tag, *auth); err != nil {
			return err
		}
	}

	return nil
}

func (s *ServiceImpl) Pull(imageRefURL string) error {
	auth, err := s.Login()
	if err != nil {
		return err
	}

	return s.dockerService.Pull(imageRefURL, *auth)
}
