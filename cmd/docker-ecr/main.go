package main

import (
	"fmt"

	"github.com/gregfurman/docker-ecr/internal/api"
	"github.com/gregfurman/docker-ecr/internal/docker"
	"github.com/gregfurman/docker-ecr/internal/ecr"
)

func main() {

	// Create clients
	// Note: perhaps look into a single factory for all clients
	// Write wrapper for clients
	ecrClient := ecr.NewClient()
	dockerClient := docker.NewClient()

	// Create new services
	dockerSvc := docker.NewService(dockerClient)
	ecrSvc := ecr.NewService(ecrClient)
	api := api.NewService(dockerSvc, ecrSvc)

	err := api.Build("Dockerfile", true, "test/repo", map[string]string{"deployer": "Greg"}, "build")
	fmt.Printf("%v\n", err)
}
