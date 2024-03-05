//nolint:gonoglobals,gochecknoglobals
package dockecr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes a docker image tagged with a repository name to a cloud repository.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := API.Push(repositoryName, repositoryTags, imageTags...); err != nil {
			log.Printf("error: %v", err)

			return
		}

		log.Println("Image pushed successfully")
	},
}

//nolint:gochecknoinits
func init() {
	pushCmd.PersistentFlags().StringVarP(&repositoryName, "repository-name", "r", "", "Repository of image")
	pushCmd.PersistentFlags().StringToStringVarP(&repositoryTags, "repository-tags", "t", map[string]string{}, "Repository resource tags to be assigned")
	pushCmd.PersistentFlags().StringSliceVarP(&imageTags, "image-tags", "i", []string{}, "docker tags to be assigned to image")

	rootCmd.AddCommand(pushCmd)
}
