package middlewares

import (
	"assesment/internal/transaction/api/rest/payloads"
	"assesment/internal/transaction/core/ports"
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func New(accountApp ports.AccountApp) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get(fiber.HeaderAuthorization)
		response := new(payloads.ResponseDefault)

		if len(auth) <= 6 || !utils.EqualFold(auth[:6], "basic ") {
			ctx.SendStatus(400)
			response.Message = "unauthorized"
			return ctx.JSON(response)
		}

		raw, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			ctx.SendStatus(400)
			response.Message = "unauthorized"
			return ctx.JSON(response)
		}

		creds := utils.UnsafeString(raw)
		index := strings.Index(creds, ":")
		if index == -1 {
			ctx.SendStatus(400)
			response.Message = "unauthorized"
			return ctx.JSON(response)
		}

		accountNumber := creds[:index]
		pin := creds[index+1:]

		if err := accountApp.CheckAccount(ctx.UserContext(), accountNumber, pin); err != nil {
			ctx.SendStatus(response.ParseFromException(err))
			return ctx.JSON(response)
		}

		return ctx.Next()
	}
}
