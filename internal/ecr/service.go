package ecr

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	CreateEcrRepository(repositoryName string, isMutableImageTags bool, repositoryTags map[string]string) (*types.Repository, error)
	GetRepositories(filter func(types.Repository) bool) ([]types.Repository, error)
	GetRepository(repositoryName string) (*types.Repository, error)
	GetAllRepositories() ([]types.Repository, error)
	GetRepositoryNamesByPrefix(prefix string) ([]string, error)
	GetImages(repositoryName string) ([]types.ImageDetail, error)
	ListImages(repositoryName string, tagStatus types.TagStatus, filter func(types.ImageDetail) bool) ([]types.ImageDetail, error)
	GetImageScanResults(repositoryName, imageDigest, imageTag string) ([]types.ImageScanFindings, error)
	GetAuth() (*types.AuthorizationData, error)
}

type ServiceImpl struct {
	client *ecr.Client
}

const maxResultsPerPage int32 = 1000

func NewService(client *ecr.Client) Service {
	return &ServiceImpl{client: client}

}

func (s *ServiceImpl) CreateEcrRepository(repositoryName string, isMutableImageTags bool, repositoryTags map[string]string) (*types.Repository, error) {

	imageMutability := types.ImageTagMutabilityMutable
	if !isMutableImageTags {
		imageMutability = types.ImageTagMutabilityImmutable
	}

	var tags []types.Tag
	for key, value := range repositoryTags {
		tags = append(tags, types.Tag{Key: &key, Value: &value})
	}

	input := ecr.CreateRepositoryInput{
		RepositoryName:     &repositoryName,
		Tags:               tags,
		ImageTagMutability: imageMutability,
	}

	out, err := s.client.CreateRepository(context.TODO(), &input)
	if err != nil {
		var dne *types.RepositoryAlreadyExistsException
		if errors.As(err, &dne) {
			return s.GetRepository(repositoryName)
		}
		return nil, err
	}

	return out.Repository, nil
}

func (s *ServiceImpl) GetRepositories(filter func(types.Repository) bool) ([]types.Repository, error) {
	maxResuls := maxResultsPerPage
	repositoriesInput := ecr.DescribeRepositoriesInput{
		MaxResults: &maxResuls,
	}

	var repositories []types.Repository
	paginator := ecr.NewDescribeRepositoriesPaginator(s.client, &repositoriesInput)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return repositories, err
		}

		for _, repo := range page.Repositories {
			if filter(repo) {
				repositories = append(repositories, repo)
			}
		}
	}

	return repositories, nil
}

func (s *ServiceImpl) GetRepository(repositoryName string) (*types.Repository, error) {
	reduce := func(repo types.Repository) bool {
		return *repo.RepositoryName == repositoryName
	}
	repo, err := s.GetRepositories(reduce)
	if err != nil {
		return nil, err
	}

	var dne *types.RepositoryNotFoundException
	if len(repo) == 0 {
		return nil, dne
	}

	return &repo[0], nil

}

func (s *ServiceImpl) GetAllRepositories() ([]types.Repository, error) {
	repos, err := s.GetRepositories(func(types.Repository) bool { return true })
	if err != nil {
		return repos, err
	}

	return repos, nil
}

func (s *ServiceImpl) GetRepositoryNamesByPrefix(prefix string) ([]string, error) {
	filter := func(repo types.Repository) bool {
		return strings.HasPrefix(*repo.RepositoryName, prefix)
	}

	repositories, err := s.GetRepositories(filter)
	if err != nil {
		return nil, err
	}

	repositoryNames := make([]string, len(repositories))
	for i, repo := range repositories {
		repositoryNames[i] = *repo.RepositoryName
	}

	return repositoryNames, nil
}

func (s *ServiceImpl) GetImages(repositoryName string) ([]types.ImageDetail, error) {
	filter := func(image types.ImageDetail) bool { return true } // return all imageDetails

	imageDetails, err := s.ListImages(repositoryName, types.TagStatusAny, filter)
	for err != nil {
		return []types.ImageDetail{}, err
	}

	return imageDetails, nil
}

func (s *ServiceImpl) ListImages(repositoryName string, tagStatus types.TagStatus, filter func(types.ImageDetail) bool) ([]types.ImageDetail, error) {

	maxResuls := maxResultsPerPage
	imagesInput := ecr.DescribeImagesInput{
		RepositoryName: &repositoryName,
		Filter:         &types.DescribeImagesFilter{TagStatus: tagStatus},
		MaxResults:     &maxResuls,
	}

	var imageDetails []types.ImageDetail
	paginator := ecr.NewDescribeImagesPaginator(s.client, &imagesInput)
	for page, err := paginator.NextPage(context.TODO()); paginator.HasMorePages(); {
		if err != nil {
			return imageDetails, err
		}

		for _, image := range page.ImageDetails {
			if filter(image) {
				imageDetails = append(imageDetails, image)
			}
		}
	}

	return imageDetails, nil
}

func (s *ServiceImpl) GetImageScanResults(repositoryName, imageDigest, imageTag string) ([]types.ImageScanFindings, error) {

	maxResults := maxResultsPerPage
	scanResultsInput := ecr.DescribeImageScanFindingsInput{
		ImageId:        &types.ImageIdentifier{ImageDigest: &imageDigest, ImageTag: &imageTag},
		RepositoryName: &repositoryName,
		MaxResults:     &maxResults,
	}

	waiter := ecr.NewImageScanCompleteWaiter(s.client)
	err := waiter.Wait(context.TODO(), &scanResultsInput, time.Second*60)
	if err != nil {
		log.Warnf("Scan waiting timed out: %v", err)
		return []types.ImageScanFindings{}, err
	}

	var scanFindings []types.ImageScanFindings
	paginator := ecr.NewDescribeImageScanFindingsPaginator(s.client, &scanResultsInput)
	for results, err := paginator.NextPage(context.TODO()); paginator.HasMorePages(); {
		if err != nil {
			return scanFindings, err
		}

		scanFindings = append(scanFindings, *results.ImageScanFindings)
	}

	return scanFindings, nil
}

func (s *ServiceImpl) GetAuth() (*types.AuthorizationData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	input := ecr.GetAuthorizationTokenInput{}

	out, err := s.client.GetAuthorizationToken(ctx, &input)
	if err != nil {
		return nil, err
	}

	if len(out.AuthorizationData) < 1 {
		return nil, fmt.Errorf("error: failed to retrieve authorisation data for ecr-login")
	}

	return &out.AuthorizationData[0], nil
}
