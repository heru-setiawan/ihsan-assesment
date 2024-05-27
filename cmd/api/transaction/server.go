package main

import (
	"assesment/internal/transaction/api/rest"
	"assesment/internal/transaction/api/rest/routes"
	"assesment/internal/transaction/app"
	"assesment/internal/transaction/core/models"
	"assesment/internal/transaction/datastore"
	"assesment/pkg/configs"
	"assesment/pkg/databases"
	"assesment/pkg/hash"
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

	log := logs.NewLogger(config.LogLogrus, "transaction")

	traceProvider := traces.NewHTTPTelemetryProvider(
		fmt.Sprintf("%s:%d", config.TraceTelemetry.Host, config.TraceTelemetry.Port),
		"transaction",
		context.Background(),
	)
	tracer := traceProvider.Tracer("transaction")

	postgresDB, err := databases.NewPostgres(config.DatabasePostgres, traceProvider)
	if err != nil {
		log.Fatal(map[string]any{"error": err}, "error connecting to postgres database")
		os.Exit(2)
	}

	if err := postgresDB.Migrate(&models.Account{}); err != nil {
		log.Fatal(map[string]any{"error": err}, "error migrating postgres database")
		os.Exit(3)
	}

	redisDB, err := databases.NewRedis(config.DatabaseRedis, traceProvider)
	if err != nil {
		log.Fatal(map[string]any{"error": err}, "error connecting to redis database")
		os.Exit(2)
	}

	hash := hash.NewBcrypt(config.HashBcrypt)

	datastore := datastore.New(*log, postgresDB.Conn, redisDB.Conn)
	app := app.New(*log, tracer, *hash, datastore)

	server := rest.NewApi(*config, *log, traceProvider, routes.Route{
		AccountApp: app,
	})

	server.Start()
}
