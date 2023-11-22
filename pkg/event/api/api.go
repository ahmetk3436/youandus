package api

import (
	"youandus/pkg/event/model"
	"youandus/pkg/event/service"

	"github.com/gofiber/fiber/v2"
)

type Event struct {
	service *service.Service
}

func NewApi(service *service.Service) Event {
	return Event{
		service: service,
	}
}

func (e Event) GetEvent(c *fiber.Ctx) error {
	eventID := c.QueryInt("eventID", 0)
	if eventID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! ",
		})
	}
	eventData, err := e.service.GetEvent(uint(eventID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Etkinlik başarıyla getirildi!",
		"data":    eventData,
	})
}

func (e Event) CreateEvent(c *fiber.Ctx) error {
	var event *model.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error when parsing body !",
		})
	}
	err := e.service.CreateEvent(event)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! " + err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Etkinlik başarıyla oluşturuldu!",
	})
}

func (e Event) UpdateEvent(c *fiber.Ctx) error {
	eventID := c.QueryInt("eventID", 0)
	if eventID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! ",
		})
	}

	var eventData model.Event

	if err := c.BodyParser(&eventData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! Geçersiz etkinlik verisi.",
		})
	}

	eventData, _ = e.service.UpdateEvent(uint(eventID), eventData)

	return c.Status(200).JSON(fiber.Map{
		"message": "Etkinlik başarıyla güncellendi!",
		"data":    eventData,
	})
}

func (e Event) DeleteEvent(c *fiber.Ctx) error {
	eventID := c.QueryInt("eventID", 0)
	if eventID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut! ",
		})
	}

	err := e.service.DeleteEvent(uint(eventID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Hata mevcut ! " + err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Etkinlik başarıyla silindi !",
	})
}
