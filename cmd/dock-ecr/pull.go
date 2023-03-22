package dockecr

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls a docker image using its resource URI.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := API.Pull(args[0]); err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		fmt.Printf("Image pulled successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
