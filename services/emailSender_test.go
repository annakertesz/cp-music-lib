package services

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/config"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/annakertesz/cp-music-lib/services/smtpClient"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"testing"
)

func TestEmailSender(t *testing.T) {
	SENDGRID_API_KEY := "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"

	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func TestEmailSender_SendSongBuyEmail(t *testing.T) {
	smtpConfig := config.SmtpEmailConfig{
		ServerAddress: "smtp.gmail.com:587",
		UserName:      "centralpublishingco@gmail.com",
		Password:      "centraladm1n",
		Host:          "smtp.gmail.com",
		ToEmail:       "forianszm@gmail.com",
		FromEmail:     "centralpublishingco@gmail.com",
	}
	sender := smtpClient.NewEmailSender(smtpConfig)
	buySongObj := models.BuySongObj{
		Title:  "test song title",
		Artist: "test artist name",
		Album:  "album name",
		User: models.UserRO{
			ID:        0,
			Username:  "testUserName",
			FirstName: "Joe",
			LastName:  "smith",
			Email:     "randomUser@example.com",
			Phone:     "00 12 3456789",
		},
		Message: "test message",
	}
	err := sender.SendSongBuyEmail(buySongObj)
	if err != nil {
		panic(err)
	}
}
