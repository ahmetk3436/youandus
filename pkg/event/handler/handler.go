package handler

import (
	"youandus/internal/storage"
	"youandus/pkg/event/api"
	"youandus/pkg/event/repository"
	"youandus/pkg/event/service"
	"youandus/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitEvent(fiber *fiber.App) {
	db := storage.GetDB()
	repo := repository.NewRepository(db)
	service := service.NewService(&repo)
	api := api.NewApi(&service)
	event := fiber.Group("/event")
	event.Use(middleware.Auth)
	event.Post("", api.CreateEvent)
	event.Get("", api.GetEvent)
	event.Put("", api.UpdateEvent)
	event.Delete("", api.DeleteEvent)
	fiber.Get("/events", api.GetEvents)
}
