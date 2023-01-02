package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapFriendsRoutes(router fiber.Router, h *FriendHandler) {
	router.Post("/friends/add", h.DeleteFriend())
	router.Post("/friends/delete", h.AddFriend())
	router.Get("/friends", h.GetFriends())
}
