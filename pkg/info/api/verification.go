package api

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"youandus/pkg/info/repository"
)

type VerificationAPI struct {
	WebApiFramework  *fiber.App
	VerificationRepo *repository.VerificationRepo
}

func NewVerificationAPI(webApiFramework *fiber.App, verfiicationRepo *repository.VerificationRepo) VerificationAPI {
	return VerificationAPI{
		WebApiFramework:  webApiFramework,
		VerificationRepo: verfiicationRepo,
	}
}
func (a VerificationAPI) HandleVerification(c *fiber.Ctx) error {
	verificationCode := c.Query("verification_code")
	email := c.Query("email")
	verificationStatus, userID, err := a.VerificationRepo.CheckEmailVerificationCode(email, verificationCode)
	if err != nil {
		errorResponse := struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
	}
	if !verificationStatus {
		return c.Redirect("http://193.164.7.10:1323/verification/dogrulama-basarisiz.html")
	}
	if verificationStatus {
		var rq struct {
			ID uint `json:"id"`
		}
		rq.ID = userID
		jsonData, err := json.Marshal(rq)
		if err != nil {
			return err
		}
		response, err := http.Post("http://193.164.7.10:1323/profile", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

	}
	return c.Redirect("http://193.164.7.10:1323/verification/dogrulama-basarili.html")
}
