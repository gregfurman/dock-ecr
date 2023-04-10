package ecr

import (
	"context"
	"errors"
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
	TagImage(repositoryName, imageDigest, imageTag string) error
}

type ServiceImpl struct {
	client *ecr.Client
}

const (
	maxResultsPerPage int32 = 1000
	requestWaitTime   int32 = 120
)

func NewService(client *ecr.Client) Service {
	return &ServiceImpl{client: client}
}

// CreateEcrRepository creates a new ECR repository called `repositoryName`, assigns it optional tags,
// and sets whether the tags of images within the repository are mutable. If successful, a pointer to a repository
// struct is returned.
func (s *ServiceImpl) CreateEcrRepository(repositoryName string, isMutableImageTags bool, repositoryTags map[string]string) (*types.Repository, error) {
	imageMutability := types.ImageTagMutabilityMutable
	if !isMutableImageTags {
		imageMutability = types.ImageTagMutabilityImmutable
	}

	tags := make([]types.Tag, len(repositoryTags))
	counter := 0

	for key := range repositoryTags {
		keyCopy := key
		value := repositoryTags[keyCopy]
		tags[counter] = types.Tag{Key: &keyCopy, Value: &value}
		counter++
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

// GetRepositories returns a subset of all repositories in the registry, filtering results based on
// the function passed in.
func (s *ServiceImpl) GetRepositories(filter func(types.Repository) bool) ([]types.Repository, error) {
	var repositories []types.Repository

	maxResuls := maxResultsPerPage
	repositoriesInput := ecr.DescribeRepositoriesInput{
		MaxResults: &maxResuls,
	}

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

// GetRepository matches and retrieves a single repostitory struct based on a provided repository name. If no
// repository found, a RepositoryNotFoundException error is returned.
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

// GetAllRepositories returns a list of all repository structs in the ECR registry.
func (s *ServiceImpl) GetAllRepositories() ([]types.Repository, error) {
	repos, err := s.GetRepositories(func(types.Repository) bool { return true })
	if err != nil {
		return repos, err
	}

	return repos, nil
}

// GetRepositoryNamesByPrefix returns the names of all repositories that match the provided prefix.
// i.e a prefix of "repository" would match the repository 012345678910.dkr.ecr.region.amazonaws.com/*repository*/name.
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

// GetImages returns a list of all images in a given repository.
func (s *ServiceImpl) GetImages(repositoryName string) ([]types.ImageDetail, error) {
	filter := func(image types.ImageDetail) bool { return true }

	imageDetails, err := s.ListImages(repositoryName, types.TagStatusAny, filter)
	for err != nil {
		return []types.ImageDetail{}, err
	}

	return imageDetails, nil
}

// ListImages returns a subset of all images in a repository, filtering by a custom passed in function as well as by TagStatus.
// TagStatus can be TagStatusTagged, TagStatusUntagged, or TagStatusAny.
func (s *ServiceImpl) ListImages(repositoryName string, tagStatus types.TagStatus, filter func(types.ImageDetail) bool) ([]types.ImageDetail, error) {
	var imageDetails []types.ImageDetail

	maxResuls := maxResultsPerPage
	imagesInput := ecr.DescribeImagesInput{
		RepositoryName: &repositoryName,
		Filter:         &types.DescribeImagesFilter{TagStatus: tagStatus},
		MaxResults:     &maxResuls,
	}

	paginator := ecr.NewDescribeImagesPaginator(s.client, &imagesInput)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
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

// GetImageScanResults returns a list of scan results for image residing in a specific repository by the image's digest and tag.
// This function will timeout and return an error if the scan has not completed in 120 seconds.
func (s *ServiceImpl) GetImageScanResults(repositoryName, imageDigest, imageTag string) ([]types.ImageScanFindings, error) {
	var scanFindings []types.ImageScanFindings

	maxResults := maxResultsPerPage
	scanResultsInput := ecr.DescribeImageScanFindingsInput{
		ImageId:        &types.ImageIdentifier{ImageDigest: &imageDigest, ImageTag: &imageTag},
		RepositoryName: &repositoryName,
		MaxResults:     &maxResults,
	}

	waiter := ecr.NewImageScanCompleteWaiter(s.client)
	if err := waiter.Wait(context.TODO(), &scanResultsInput, time.Second*time.Duration(requestWaitTime)); err != nil {
		log.Warnf("Scan waiting timed out: %v", err)

		return []types.ImageScanFindings{}, err
	}

	paginator := ecr.NewDescribeImageScanFindingsPaginator(s.client, &scanResultsInput)

	for paginator.HasMorePages() {
		results, err := paginator.NextPage(context.TODO())
		if err != nil {
			return scanFindings, err
		}

		scanFindings = append(scanFindings, *results.ImageScanFindings)
	}

	return scanFindings, nil
}

// GetAuth generates a struct with authorisation credentials required to interface with ECR.
func (s *ServiceImpl) GetAuth() (*types.AuthorizationData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(requestWaitTime))
	defer cancel()

	input := ecr.GetAuthorizationTokenInput{}

	out, err := s.client.GetAuthorizationToken(ctx, &input)
	if err != nil {
		return nil, err
	}

	if len(out.AuthorizationData) < 1 {
		return nil, errNoAuthData
	}

	return &out.AuthorizationData[0], nil
}

func (s *ServiceImpl) TagImage(repositoryName, imageDigest, imageTag string) error {
	input := ecr.PutImageInput{
		RepositoryName: &repositoryName,
		ImageDigest:    &imageDigest,
		ImageTag:       &imageTag,
	}

	_, err := s.client.PutImage(context.TODO(), &input)

	return err
}
