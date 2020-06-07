package controller

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	SENDGRID_API_KEY = "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"
	SENDER_NAME = "Central Publishing"
	SENDER_EMAIL = "kerteszannanak@gmail.com"
	CENTRAL_ADMIN_ADRESS = "centralpublishingtest@gmail.com"
	)

func sendVerifyEmail(user models.UserReqObj, verifyEndpoint string) error{
	from := mail.NewEmail(SENDER_NAME, SENDER_EMAIL)
	subject := fmt.Sprintf("%v %v would like have access", user.FirstName, user.LastName)
	to := mail.NewEmail("Central admin", CENTRAL_ADMIN_ADRESS)
	plainTextContent := "."
	htmlContent := fmt.Sprintf(
		"<h2 style=\"color: #2e6c80;\">New user request:</h2>" +
			"<ul>" +
			"<li>Name: %v %v</li>" +
			"<li>Email: %v</li>" +
			"<li>Phone number: %v</li>" +
			"</ul>" +
			"<p>Please click <a href=\"http://%v\">here</a> to accept his/her request</p>",
			user.FirstName, user.LastName, user.Email, user.Phone, verifyEndpoint)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	return sendEmail(message)
}

func sendEmail(message *mail.SGMailV3) error {
	client := sendgrid.NewSendClient(SENDGRID_API_KEY)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}