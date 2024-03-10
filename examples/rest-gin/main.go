package main

import (
	"context"
	"log/slog"

	"github.com/livraghi/microservice"
	"github.com/livraghi/microservice/configuration"
)

var (
	BuildVersion = "(devel)"
)

func main() {
	ctx := context.TODO()

	cfg, err := configuration.LoadConfigurations(
		configuration.WithAppVersion(BuildVersion),
		configuration.WithConfigPath("./configs"),
		configuration.WithConfigName("local"),
		configuration.WithConfigType(configuration.ENV),
	)
	if err != nil {
		slog.Error("failed to load configurations", "error", err.Error())
		return
	}

	microservice.StartDefaultHttpServer(ctx, microservice.WithConfiguration(cfg))
}
