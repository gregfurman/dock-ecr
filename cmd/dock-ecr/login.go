//nolint:gonoglobals,gochecknoglobals
package dockecr

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs into an ECR registry.",
	Run: func(cmd *cobra.Command, args []string) {
		auth64, err := API.Login()
		if err != nil {
			log.Printf("error: %v", err)

			return
		}

		log.Println("Login successful")

		auth, _ := base64.URLEncoding.DecodeString(*auth64)

		//nolint:tagliatelle
		authConfig := struct {
			Username string `json:"Username"`
			Password string `json:"Password"`
		}{}

		if err := json.Unmarshal(auth, &authConfig); err != nil {
			log.Printf("error: %v", err)
		}

		fmt.Print(authConfig.Password)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(loginCmd)
}
