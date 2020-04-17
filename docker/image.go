package docker

import (
	"bufio"
	"context"
	"encoding/json"
	"strings"

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
		log.Fatal(err)
	}
	// read docker response
	if verbose {
		scanner := bufio.NewScanner(response.Body)
		for scanner.Scan() {
			bytes := []byte(scanner.Text())
			data := make(map[string]interface{})
			if err := json.Unmarshal(bytes, &data); err != nil {
				log.Errorln("Can't process docker output")
			}
			// just get strings
			if line, ok := data["stream"].(string); ok {
				line := strings.ReplaceAll(line, "\n", "")
				if line != "" {
					log.Info(" > ", line)
				}
			}
		}
	}
	response.Body.Close()

	return
}
