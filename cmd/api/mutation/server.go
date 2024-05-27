package main

import (
	"assesment/internal/mutation/api/stream"
	"assesment/internal/mutation/app"
	"assesment/internal/mutation/core/models"
	"assesment/internal/mutation/datastore"
	"assesment/pkg/configs"
	"assesment/pkg/databases"
	"assesment/pkg/logs"
	"assesment/pkg/traces"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	config := new(configs.Config)
	err := config.LoadFromEnvFile()
	if err != nil {
		log.Println("error loading config: ", err)
		os.Exit(1)
	}

	log := logs.NewLogger(config.LogLogrus, "mutation")

	traceProvider := traces.NewHTTPTelemetryProvider(
		fmt.Sprintf("%s:%d", config.TraceTelemetry.Host, config.TraceTelemetry.Port),
		"mutation",
		context.Background(),
	)
	tracer := traceProvider.Tracer("mutation")

	postgresDB, err := databases.NewPostgres(config.DatabasePostgres, traceProvider)
	if err != nil {
		log.Fatal(map[string]any{"error": err}, "error connecting to postgres database")
		os.Exit(2)
	}

	if err := postgresDB.Migrate(&models.Journal{}); err != nil {
		log.Fatal(map[string]any{"error": err}, "error migrating postgres database")
		os.Exit(3)
	}

	redisDB, err := databases.NewRedis(config.DatabaseRedis, traceProvider)
	if err != nil {
		log.Fatal(map[string]any{"error": err}, "error connecting to redis database")
		os.Exit(2)
	}

	datastore := datastore.New(*log, postgresDB.Conn)
	app := app.New(*log, tracer, datastore)

	server, err := stream.NewRedis(redisDB.Conn, app)
	if err != nil {
		log.Fatal(map[string]any{"error": err}, "error initializing redis stream")
		os.Exit(4)
	}
	server.Start()
}
