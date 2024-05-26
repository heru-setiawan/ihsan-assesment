package routes

import (
	"assesment/internal/api/rest/payloads"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Tabung(c *fiber.Ctx) error {
	var request = new(payloads.RequestTransaksiRekening)
	var response = new(payloads.ResponseDefault)

	c.BodyParser(&request)

	result, err := r.RekeningApp.Tabung(c.Context(), request.No, request.Nominal)
	if err != nil {
		c.SendStatus(response.ParseFromException(err))
		return c.JSON(response)
	}

	c.SendStatus(200)
	response.Data = result
	return c.JSON(response)
}

func (r *Route) Tarik(c *fiber.Ctx) error {
	var request = new(payloads.RequestTransaksiRekening)
	var response = new(payloads.ResponseDefault)

	c.BodyParser(&request)

	result, err := r.RekeningApp.Tarik(c.Context(), request.No, request.Nominal)
	if err != nil {
		c.SendStatus(response.ParseFromException(err))
		return c.JSON(response)
	}

	c.SendStatus(200)
	response.Data = result
	return c.JSON(response)
}

func (r *Route) CekSaldo(c *fiber.Ctx) error {
	var response = new(payloads.ResponseDefault)

	noRekening := c.Params("no_rekening", "")
	result, err := r.RekeningApp.CekSaldo(c.Context(), noRekening)
	if err != nil {
		c.SendStatus(response.ParseFromException(err))
		return c.JSON(response)
	}

	c.SendStatus(200)
	response.Data = result
	return c.JSON(response)
}
