package routes

import (
	"assesment/internal/transaction/api/rest/payloads"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Register(ctx *fiber.Ctx) error {
	var request = new(payloads.RequestRegistration)
	var response = new(payloads.ResponseDefault)

	ctx.BodyParser(&request)

	result, err := r.AccountApp.Register(ctx.UserContext(), request.PIN)
	if err != nil {
		ctx.SendStatus(response.ParseFromException(err))
		return ctx.JSON(response)
	}

	ctx.SendStatus(200)
	response.Data = result
	return ctx.JSON(response)
}

func (r *Route) Deposit(ctx *fiber.Ctx) error {
	var request = new(payloads.RequestTransaction)
	var response = new(payloads.ResponseDefault)

	ctx.BodyParser(&request)

	result, err := r.AccountApp.Deposit(ctx.UserContext(), request.Number, request.Amount)
	if err != nil {
		ctx.SendStatus(response.ParseFromException(err))
		return ctx.JSON(response)
	}

	ctx.SendStatus(200)
	response.Data = result
	return ctx.JSON(response)
}

func (r *Route) Withdraw(ctx *fiber.Ctx) error {
	var request = new(payloads.RequestTransaction)
	var response = new(payloads.ResponseDefault)

	ctx.BodyParser(&request)

	result, err := r.AccountApp.Withdraw(ctx.UserContext(), request.Number, request.Amount)
	if err != nil {
		ctx.SendStatus(response.ParseFromException(err))
		return ctx.JSON(response)
	}

	ctx.SendStatus(200)
	response.Data = result
	return ctx.JSON(response)
}
