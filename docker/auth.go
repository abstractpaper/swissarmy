package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

// AuthRegistry authenticates with a docker registry.
func AuthRegistry(ctx context.Context, cli *client.Client, username string, password string, email string, server string) (err error) {
	config := types.AuthConfig{
		Username:      username,
		Password:      password,
		Email:         email,
		ServerAddress: server,
	}
	body, err := cli.RegistryLogin(ctx, config)
	if err != nil {
		log.Errorln("Error in authenticating with docker registry: ", err)
	}
	log.Infoln("RegistryLogin response: ", body.Status)

	return
}
