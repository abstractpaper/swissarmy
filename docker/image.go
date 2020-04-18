package docker

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"

	swissJSON "github.com/abstractpaper/swissarmy/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	log "github.com/sirupsen/logrus"
)

// BuildImage builds a docker image residing in path.
func BuildImage(ctx context.Context, cli *client.Client, path string, dockerFile string, tags []string, verbose bool) (err error) {
	// tar file
	buildContext, err := archive.Tar(path, archive.Uncompressed)

	// default dockerfile
	if dockerFile == "" {
		dockerFile = "Dockerfile"
	}
	// build options
	opt := types.ImageBuildOptions{
		Dockerfile: dockerFile,
		Tags:       tags,
	}
	// build image
	response, err := cli.ImageBuild(ctx, buildContext, opt)
	if err != nil {
		log.Error(err)
	}
	// read docker response
	if verbose {
		parseOutput(response.Body)
	}
	response.Body.Close()

	return
}

// PushImage pushes a docker image to a registry.
// Use full URLs for private registries (e.g. AWS ECR)
func PushImage(ctx context.Context, cli *client.Client, image string, username string, password string, email string, verbose bool) (err error) {
	bytes, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	})

	token := base64.StdEncoding.EncodeToString(bytes)
	out, err := cli.ImagePush(ctx, image, types.ImagePushOptions{RegistryAuth: token})
	if err != nil {
		log.Error(err)
	}
	if verbose {
		parseOutput(out)
	}
	out.Close()

	return
}

func parseOutput(body io.ReadCloser) {
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		bytes := []byte(scanner.Text())
		data := make(map[string]interface{})
		if err := json.Unmarshal(bytes, &data); err != nil {
			log.Errorln("Can't process docker output")
		}

		swissJSON.Print(data, swissJSON.PrintOptions{
			Prepend:          ">",
			SkipEmptyStrings: true,
			TrimStrings:      true,
		})
	}
}
