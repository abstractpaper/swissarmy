package docker

import (
	"context"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

// Client returns a new docker client.
func Client() (cli *client.Client, ctx context.Context, err error) {
	ctx = context.Background()
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Errorln("Error in setupDockerCLI: ", err)
	}
	cli.NegotiateAPIVersion(ctx)

	return
}
