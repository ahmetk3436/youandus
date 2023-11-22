package handler

import (
	"youandus/internal/storage"
	"youandus/pkg/profile/api"
	"youandus/pkg/profile/middleware"
	"youandus/pkg/profile/repository"
	"youandus/pkg/profile/service"

	"github.com/gofiber/fiber/v2"
)

func InitHandler(fiber *fiber.App) {
	db := storage.GetDB()
	repo := repository.NewRepository(db)
	service := service.NewService(&repo)
	api := api.NewApi(&service)
	profile := fiber.Group("/profile")
	profile.Post("", api.CreateProfile)
	profile.Use(middleware.Auth)
	profile.Get("", api.GetProfile)
	profile.Put("", api.UpdateProfile)
	profile.Delete("", api.DeleteProfile)
}
