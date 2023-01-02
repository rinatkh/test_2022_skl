package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapOrderRoutes(router fiber.Router, h *OrderHandler) {
	router.Get("/orders/:order_id", h.GetOrder())
	router.Put("/orders/:order_id/add", h.AddOrderProducts())
	router.Put("/orders/:order_id/delete", h.DeleteOrderProducts())
	router.Delete("/orders/:order_id", h.DeleteOrder())
	router.Get("/orders", h.GetOrders())
	router.Post("/orders", h.CreateOrder())
}
