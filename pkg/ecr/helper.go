package ecr

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type authConfig struct {
	Username string
	Password string
}

func FormatAuthDetails(auth string) (*string, error) {
	decodedAuthToken, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(decodedAuthToken), ":")

	if len(parts) != 2 {
		return nil, errors.New("decoded authenication string is not in the format 'username:password'")
	}

	authcfg := authConfig{Username: parts[0], Password: parts[1]}

	authConfigBytes, _ := json.Marshal(authcfg)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	return &authConfigEncoded, nil

}
