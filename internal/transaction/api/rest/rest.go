package rest

import (
	"assesment/internal/transaction/api/rest/middlewares"
	"assesment/internal/transaction/api/rest/routes"
	"assesment/pkg/configs"
	"assesment/pkg/logs"
	"fmt"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.opentelemetry.io/otel/trace"
)

func NewApi(config configs.Config, log logs.Logger, tracerProvider trace.TracerProvider, router routes.Route) *restApi {
	var api = new(restApi)
	api.config = config
	api.log = log
	api.router = router
	api.tracerProvider = tracerProvider

	api.init()
	api.registerMiddleware()
	api.registerRoute()

	return api
}

type restApi struct {
	config         configs.Config
	Server         *fiber.App
	log            logs.Logger
	tracerProvider trace.TracerProvider
	router         routes.Route
}

func (a *restApi) init() {
	a.Server = fiber.New()
}

func (a *restApi) registerMiddleware() {
	a.Server.Use(recover.New())
	a.Server.Use(cors.New())
	a.Server.Use(otelfiber.Middleware(otelfiber.WithTracerProvider(a.tracerProvider)))
}

func (a *restApi) registerRoute() {
	a.Server.Post("/daftar", a.router.Register)

	authGroup := a.Server.Group("/", middlewares.New(a.router.AccountApp))
	authGroup.Post("/tabung", a.router.Deposit, middlewares.New(a.router.AccountApp))
	authGroup.Post("/tarik", a.router.Withdraw, middlewares.New(a.router.AccountApp))
}

func (a *restApi) Start() error {
	return a.Server.Listen(fmt.Sprintf(":%d", a.config.ServiceTransaction.Port))
}
