package dockecr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes a docker image tagged with a repository name to a cloud repository.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := API.Push(repositoryName, repositoryTags); err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		fmt.Printf("Image pushed successfully\n")
	},
}

func init() {
	pushCmd.PersistentFlags().StringVarP(&repositoryName, "repository-name", "r", "", "Repository of image")
	pushCmd.PersistentFlags().StringToStringVarP(&repositoryTags, "repository-tags", "t", map[string]string{}, "Repository resource tags to be assigned")

	rootCmd.AddCommand(pushCmd)
}
