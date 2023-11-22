package api

import (
	"youandus/helper"
	"youandus/pkg/profile/model"
	"youandus/pkg/profile/service"

	"github.com/gofiber/fiber/v2"
)

type Api struct {
	service *service.Service
}

func NewApi(service *service.Service) Api {
	return Api{
		service: service,
	}
}
func (a Api) GetProfile(c *fiber.Ctx) error {
	userID, err := helper.GetUserID(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})
	}
	profile, err := a.service.GetProfile(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Profil başarıyla getirildi!",
		"data":    profile,
	})
}
func (a Api) CreateProfile(c *fiber.Ctx) error {
	var rq struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(&rq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error when parsing body !",
		})
	}
	err := a.service.CreateProfile(rq.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})

	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Profil başarıyla oluşturuldu!",
	})
}

func (a Api) UpdateProfile(c *fiber.Ctx) error {
	userID, err := helper.GetUserID(c)
	println(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})

	}

	var profile model.User

	// Kullanıcıdan gelen JSON verisini profile değişkenine bind etmek
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! Geçersiz profil verisi.",
		})

	}

	profile, err = a.service.UpdateProfile(userID, profile)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})

	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Profil başarıyla güncellendi!",
		"data":    profile,
	})
}

func (a Api) DeleteProfile(c *fiber.Ctx) error {
	userID, err := helper.GetUserID(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut ! " + err.Error(),
		})
	}

	err = a.service.DeleteProfile(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut ! " + err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Profil başarıyla silindi !",
	})
}
