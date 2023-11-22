package handler

import (
	"github.com/gofiber/fiber/v2"
	"youandus/internal/storage"
	"youandus/pkg/middleware"
	"youandus/pkg/profile/api"
	"youandus/pkg/profile/repository"
	"youandus/pkg/profile/service"
)

func InitProfile(fiber *fiber.App) {
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
