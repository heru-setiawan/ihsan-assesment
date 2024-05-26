package routes

import (
	"assesment/internal/api/rest/payloads"
	"assesment/internal/core/models"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Daftar(c *fiber.Ctx) error {
	var request = new(models.Nasabah)
	var response = new(payloads.ResponseDefault)

	c.BodyParser(&request)

	result, err := r.NasabahApp.Daftar(c.Context(), *request)
	if err != nil {
		c.SendStatus(response.ParseFromException(err))
		return c.JSON(response)
	}

	c.SendStatus(200)
	response.Data = result.Rekening
	return c.JSON(response)
}
