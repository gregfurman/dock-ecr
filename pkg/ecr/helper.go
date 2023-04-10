package ecr

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	errAuthFmt    = errors.New("decoded authentication string is not in the format 'username:password'")
	errNoAuthData = errors.New("failed to retrieve authorisation data for ecr-login")
)

//nolint:tagliatelle
type authConfig struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func FormatAuthDetails(auth string) (*string, error) {
	decodedAuthToken, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(decodedAuthToken), ":")

	//nolint:gomnd
	if len(parts) != 2 {
		return nil, errAuthFmt
	}

	authcfg := authConfig{Username: parts[0], Password: parts[1]}

	authConfigBytes, err := json.Marshal(authcfg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall authentication config to JSON: %w", err)
	}

	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	return &authConfigEncoded, nil
}
