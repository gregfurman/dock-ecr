//nolint:gonoglobals,gochecknoglobals
package dockecr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var image string

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls a docker image using its resource URI.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := API.Pull(image); err != nil {
			log.Printf("error: %v", err)

			return
		}

		log.Printf("Image pulled successfully\n")
	},
}

//nolint:gochecknoinits
func init() {
	pullCmd.PersistentFlags().StringVarP(&image, "image-name", "i", "", "URI of image")

	rootCmd.AddCommand(pullCmd)
}
