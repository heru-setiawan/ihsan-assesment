package main

import (
	"assesment/internal/api/rest"
	"assesment/internal/api/rest/routes"
	"assesment/internal/app"
	"assesment/internal/core/models"
	"assesment/internal/datastore"
	"assesment/pkg/configs"
	"assesment/pkg/databases"
	"assesment/pkg/logs"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var config = new(configs.Config)
	config.LoadFromEnv()

	var log = logs.NewLogger(*config)

	conn, err := databases.ConnectDB(*config)
	if err != nil {
		log.Error(map[string]any{"error": err}, "can't connect to database")
		os.Exit(1)
	}

	databases.Migrate(conn, *log, *config, &models.Nasabah{}, &models.Rekening{})

	database := datastore.NewDatabase(conn)
	app := app.New(*log, database, database)

	api := rest.NewApi(*config, *log, routes.Route{
		NasabahApp:  app,
		RekeningApp: app,
	})

	if err := api.Start(); err != nil {
		log.Error(map[string]any{"error": err}, "couldn't start service due something error")
	}
}
