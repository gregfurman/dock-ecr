package dockecr

import (
	"fmt"
	"os"

	"github.com/gregfurman/docker-ecr/internal/api"
	"github.com/gregfurman/docker-ecr/internal/docker"
	"github.com/gregfurman/docker-ecr/internal/ecr"
	"github.com/spf13/cobra"
)

var (
	API     api.Service
	rootCmd = &cobra.Command{
		Use:              "dock-ecr",
		Short:            "dock-ecr - a single CLI to interact with a local docker and a cloud provider",
		Long:             ``,
		PersistentPreRun: initAPI,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("docker-ecr tool")
		},
	}
)

func initAPI(cmd *cobra.Command, args []string) {
	// Create clients
	ecrClient := ecr.NewClient()
	dockerClient := docker.NewClient()

	// Create new services
	dockerSvc := docker.NewService(dockerClient)
	ecrSvc := ecr.NewService(ecrClient)
	API = api.NewService(dockerSvc, ecrSvc)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
