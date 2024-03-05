//nolint:gonoglobals,gochecknoglobals
package dockecr

import (
	"fmt"
	"os"

	"github.com/gregfurman/docker-ecr/pkg/api"
	"github.com/gregfurman/docker-ecr/pkg/docker"
	"github.com/gregfurman/docker-ecr/pkg/ecr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var (
	API     api.Service
	rootCmd = &cobra.Command{
		Use:              "dock-ecr",
		Short:            "dock-ecr - a single CLI to interact with a local docker server and a cloud provider",
		Long:             ``,
		Version:          Version,
		PersistentPreRun: initAPI,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("docker-ecr tool %s", Version)
		},
	}
)

func initAPI(_ *cobra.Command, _ []string) {
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
