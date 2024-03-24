package main

import (
	"context"
	"github.com/livraghi/microservice/observability"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

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
		slog.Error(
			"failed to load configurations", "error", err.Error())
		return
	}

	stopObservabilityFn, err := observability.SetUpObservability(cfg.Name, cfg.Version)
	if err != nil {
		slog.Error("failed to setup observability", "error", err.Error())
		return
	}
	defer func() { _ = stopObservabilityFn(ctx) }()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	http.Handle("/", r)

	microservice.StartDefaultHttpServer(ctx, microservice.WithConfiguration(cfg))
}
