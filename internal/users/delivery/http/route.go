package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(router fiber.Router, h *UserHandler) {
	router.Get("/users/:user_id", h.GetUser())
	router.Put("/users/:user_id", h.UpdateUser())
	router.Delete("/users/:user_id", h.DeleteUser())
	router.Get("/users", h.GetUsers())
	router.Post("/users", h.CreateUser())
}
