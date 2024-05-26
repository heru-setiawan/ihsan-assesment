package routes

import (
	"assesment/internal/core/ports"
)

type Route struct {
	NasabahApp  ports.NasabahApp
	RekeningApp ports.RekeningApp
}
