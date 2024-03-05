package docker

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
	"regexp"

	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
)

func CreateAuthConfig(username, password string) (string, error) {
	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}

	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	return authStr, nil
}

func GetDockerfileImages(loc string) ([]string, error) {
	file, err := os.Open(loc)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`^FROM\s+(?P<image>[^\s]+)`)

	var images []string
	for scanner.Scan() {
		line := scanner.Text()
		if match := pattern.FindStringSubmatch(line); match != nil {
			images = append(images, match[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return images, nil
}

func parse(rd io.Reader) error {
	type ErrorDetail struct {
		Message string `json:"message"`
	}

	type ErrorLine struct {
		Error       string      `json:"error"`
		ErrorDetail ErrorDetail `json:"errorDetail"`
	}

	type Line struct {
		Status string `json:"status"`
	}

	var lastLine, lastStatus string

	line := &Line{}

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {

		lastLine = scanner.Text()
		if err := json.Unmarshal([]byte(lastLine), line); err != nil {
			log.Warnf("Cannot unmarshall string [%s]: %v\n", lastLine, err)
		}

		if lastStatus != line.Status {
			lastStatus = line.Status
			log.Println(lastStatus)
		}
	}

	errLine := &ErrorLine{}
	if err := json.Unmarshal([]byte(lastLine), errLine); err != nil {
		log.Warnf("Cannot unmarshall string [%s]: %v\n", lastLine, err)
	}

	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
