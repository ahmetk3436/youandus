package api

import (
	"youandus/helper"
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

func (e Event) handleError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}

func (e Event) GetEvents(c *fiber.Ctx) error {
	userID, err := helper.GetUserID(c)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Kullanıcı bilgisi alınamadı: "+err.Error())
	}

	eventData, err := e.service.GetEvents(userID)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik getirilirken hata oluştu: "+err.Error())
	}
	if eventData == nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik verisi boş ! : "+err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Etkinlikler başarıyla getirildi!",
		"data":    eventData,
	})
}

func (e Event) GetEvent(c *fiber.Ctx) error {
	eventID := c.QueryInt("eventID", 0)
	if eventID == 0 {
		return e.handleBadRequest(c, "Geçersiz etkinlik kimliği!")
	}

	userID, err := helper.GetUserID(c)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Kullanıcı bilgisi alınamadı: "+err.Error())
	}

	eventData, err := e.service.GetEvent(uint(eventID), userID)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik getirilirken hata oluştu: "+err.Error())
	}
	if eventData == nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik verisi boş ! : "+err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Etkinlik başarıyla getirildi!",
		"data":    eventData,
	})
}

func (e Event) handleBadRequest(c *fiber.Ctx, message string) error {
	return e.handleError(c, fiber.StatusBadRequest, message)
}

func (e Event) CreateEvent(c *fiber.Ctx) error {
	var event *model.Event
	if err := c.BodyParser(&event); err != nil {
		return e.handleBadRequest(c, "Body parse hatası: "+err.Error())
	}

	userID, err := helper.GetUserID(c)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Kullanıcı bilgisi alınamadı: "+err.Error())
	}

	if isEmptyEvent(event) {
		return e.handleBadRequest(c, "Etkinlik bilgileri eksik!")
	}

	event.UserID = userID
	err = e.service.CreateEvent(event)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik oluşturulurken hata oluştu: "+err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Etkinlik başarıyla oluşturuldu!",
	})
}

func (e Event) UpdateEvent(c *fiber.Ctx) error {
	eventID := c.QueryInt("eventID", 0)
	if eventID == 0 {
		return e.handleBadRequest(c, "Geçersiz etkinlik kimliği!")
	}

	var eventData *model.Event
	if err := c.BodyParser(&eventData); err != nil {
		return e.handleBadRequest(c, "Hata mevcut! Geçersiz etkinlik verisi: "+err.Error())
	}

	if isEmptyEvent(eventData) {
		return e.handleBadRequest(c, "Etkinlik bilgileri eksik!")
	}

	eventData, err := e.service.UpdateEvent(uint(eventID), *eventData)
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik güncellenirken hata oluştu: "+err.Error())
	}
	if eventData == nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik verisi boş ! : "+err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Etkinlik başarıyla güncellendi!",
		"data":    eventData,
	})
}

func (e Event) DeleteEvent(c *fiber.Ctx) error {
	eventID := c.QueryInt("eventID", 0)
	if eventID == 0 {
		return e.handleBadRequest(c, "Geçersiz etkinlik kimliği!")
	}

	err := e.service.DeleteEvent(uint(eventID))
	if err != nil {
		return e.handleError(c, fiber.StatusInternalServerError, "Etkinlik silinirken hata oluştu: "+err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Etkinlik başarıyla silindi!",
	})
}

func isEmptyEvent(event *model.Event) bool {
	return event == nil || event.EventName == "" || event.EventDescription == "" || event.EventDate == "" || event.EventLocation == "" ||
		event.EventType == "" || event.Organizer == "" || event.ContactEmail == "" || event.ContactPhone == "" ||
		event.Website == "" || event.Capacity == 0
}
