package controller

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailSender struct {
	apiKey string
	senderName string
	senderMail string
	adminEmail string
	developerEmail string
}

func NewEmailSender(apiKey string, sendername string, senderMail string, adminMail string, developerMail string) EmailSender {
	return EmailSender{
		apiKey:apiKey,
		senderName:sendername,
		senderMail:senderMail,
		adminEmail:adminMail,
		developerEmail:developerMail,
	}
}
//
//const (
//	SENDGRID_API_KEY = "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"
//	SENDER_NAME = "Central Publishing"
//	SENDER_EMAIL = "kerteszannanak@gmail.com"
//	CENTRAL_ADMIN_ADRESS = "centralpublishingtest@gmail.com"
//	)

func (sender *EmailSender) sendVerifyEmail(user models.UserReqObj, verifyEndpoint string) error{
	from := mail.NewEmail(sender.senderName, sender.senderMail)
	subject := fmt.Sprintf("%v %v would like have access", user.FirstName, user.LastName)
	to := mail.NewEmail("Central admin", sender.adminEmail)
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
	return sendEmail(message, sender.apiKey)
}

func sendEmail(message *mail.SGMailV3, apiKey string) error {
	client := sendgrid.NewSendClient(apiKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}