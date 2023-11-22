package handler

import (
	"youandus/internal/storage"
	"youandus/pkg/event/api"
	"youandus/pkg/event/middleware"
	"youandus/pkg/event/repository"
	"youandus/pkg/event/service"

	"github.com/gofiber/fiber/v2"
)

func InitHandler(fiber *fiber.App) {
	db := storage.GetDB()
	repo := repository.NewRepository(db)
	service := service.NewService(&repo)
	api := api.NewApi(&service)
	event := fiber.Group("/event")
	event.Post("", api.CreateEvent)
	event.Use(middleware.Auth)
	event.Get("", api.GetEvent)
	event.Put("", api.UpdateEvent)
	event.Delete("", api.DeleteEvent)
}
