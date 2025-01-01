package routes

import (
	"github.com/gofiber/fiber/v2"
	"hotels-service/api/http/handlers"
	"hotels-service/app"
)

func RegisterRoomRoutes(appContainer app.App, api fiber.Router) {
	roomHandler := handlers.NewRoomHandler(appContainer)
	rooms := api.Group("/rooms")
	rooms.Post("/", roomHandler.CreateRoom)
	rooms.Get("/:id", roomHandler.GetRoom)
	rooms.Get("/", roomHandler.GetAllRooms)
	rooms.Get("/available", roomHandler.GetAvailableRooms)
	rooms.Get("/search", roomHandler.FindRooms)
	rooms.Put("/:id", roomHandler.UpdateRoom)
	rooms.Put("/:id/status", roomHandler.SetRoomStatus)
	rooms.Delete("/:id", roomHandler.DeleteRoom)
}
