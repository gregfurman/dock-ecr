package api

import (
	"github.com/gregfurman/docker-ecr/internal/docker"
	"github.com/gregfurman/docker-ecr/internal/ecr"
)

type Service interface {
	Build(imageRefUrl string, push bool, repositoryName string, repositoryTags map[string]string, imageTags ...string) error
	Push(repositoryName string, repositoryTags map[string]string) error
	Pull(imageRefUrl string) error
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

func (c *ServiceImpl) Login() (*string, error) {
	auth, err := c.ecrService.GetAuth()
	if err != nil {
		return nil, err
	}

	return auth.AuthorizationToken, nil
}

func (s *ServiceImpl) Build(imageRefUrl string, push bool, repositoryName string, repositoryTags map[string]string, imageTags ...string) error {
	imageTags = append(imageTags, repositoryName)

	if err := s.dockerService.Build(imageRefUrl, imageTags...); err != nil {
		return err
	}

	if !push {
		return nil
	}

	if err := s.Push(repositoryName, repositoryTags); err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) Push(repositoryName string, repositoryTags map[string]string) error {
	auth, err := s.Login()
	if err != nil {
		return err
	}

	repo, err := s.ecrService.CreateEcrRepository(repositoryName, false, repositoryTags)
	if err != nil {
		return err
	}

	if err := s.dockerService.Tag(repositoryName, *repo.RepositoryUri); err != nil {
		return err
	}

	fmtAuth, err := ecr.FormatAuthDetails(*auth)
	if err != nil {
		return err
	}

	return s.dockerService.Push(*repo.RepositoryUri, *fmtAuth)
}

func (c *ServiceImpl) Pull(imageRefUrl string) error {
	authorisationToken, err := c.Login()
	if err != nil {
		return err
	}

	return c.dockerService.Pull(imageRefUrl, *authorisationToken)
}
