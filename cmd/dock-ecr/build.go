package dockecr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	dockerfile, repositoryName string
	imageTags                  []string
	mustPush                   bool
	repositoryTags             map[string]string
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a docker image and, if specified, pushes it to a cloud repository.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := API.Build(dockerfile, mustPush, repositoryName, repositoryTags, imageTags...); err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		fmt.Printf("Image built successfully\n")
	},
}

func init() {

	buildCmd.PersistentFlags().StringArrayVarP(&imageTags, "image-tags", "i", []string{}, "docker tags to be assigned to image")
	buildCmd.PersistentFlags().StringVarP(&dockerfile, "dockerfile", "d", "Dockerfile", "Path to Dockerfile")
	buildCmd.PersistentFlags().StringVarP(&repositoryName, "repository-name", "r", "", "Repository of image")
	buildCmd.PersistentFlags().BoolVar(&mustPush, "push", false, "If `true`, pushes the image to the specified repository")
	buildCmd.PersistentFlags().StringToStringVarP(&repositoryTags, "repository-tags", "t", map[string]string{}, "Repository resource tags to be assigned")

	rootCmd.AddCommand(buildCmd)
}
