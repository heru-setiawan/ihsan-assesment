package rest

import (
	"assesment/internal/api/rest/routes"
	"assesment/pkg/configs"
	"assesment/pkg/logs"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApi(config configs.Config, log logs.Logger, router routes.Route) *restApi {
	var api = new(restApi)
	api.config = config
	api.log = log
	api.router = router

	api.init()
	api.registerMiddleware()
	api.registerRoute()

	return api
}

type restApi struct {
	config configs.Config
	Server *fiber.App
	log    logs.Logger
	router routes.Route
}

func (a *restApi) init() {
	a.Server = fiber.New()
}

func (a *restApi) registerMiddleware() {
	a.Server.Use(recover.New())
	a.Server.Use(cors.New())
}

func (a *restApi) registerRoute() {
	a.Server.Post("/daftar", a.router.Daftar)
	a.Server.Post("/tabung", a.router.Tabung)
	a.Server.Post("/tarik", a.router.Tarik)
	a.Server.Get("/saldo/:no_rekening", a.router.CekSaldo)
}

func (a *restApi) Start() error {
	return a.Server.Listen(fmt.Sprintf(":%d", a.config.Service.Port))
}
