package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapProductRoutes(router fiber.Router, h *ProductHandler) {
	router.Get("/products/:product_id", h.GetProduct())
	router.Put("/products/:product_id", h.UpdateProduct())
	router.Delete("/products/:product_id", h.DeleteProduct())
	router.Get("/products", h.GetProducts())
	router.Post("/products", h.CreateProduct())
}
