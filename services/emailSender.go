package services

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/config"
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

func NewEmailSender(config config.EmailConfig) EmailSender {
	return EmailSender{
		senderName:	config.SenderName,
		senderMail: config.SenderEmail,
		adminEmail: config.AdminEmail,
		developerEmail: config.DeveloperEmail,
	}
}

func (sender *EmailSender) SendVerifyEmail(user models.UserReqObj, verifyEndpoint string) error{
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

func (sender *EmailSender) SendSongBuyEmail(msg models.BuySongObj) error{
	from := mail.NewEmail(sender.senderName, sender.senderMail)
	subject := fmt.Sprintf("%v %v would like to buy a song", msg.User.FirstName, msg.User.LastName)
	to := mail.NewEmail("Central admin", sender.adminEmail)
	plainTextContent := "."
	htmlContent := fmt.Sprintf(
		"<h2 style=\"color: #2e6c80;\">New song request:</h2>" +
			"<p>Dear admin</p>" +
			"<p>%v %v would like to buy a song:</p>" +
			"<p>%v</p>" +
			"<ul>" +
			"<li>Title: %v</li>" +
			"<li>Album: %v</li>" +
			"<li>Artist: %v</li>" +
			"</ul>" +
			"<p>You can contact him/her via email (%v) or phone (%v)</p>",
		msg.User.FirstName, msg.User.LastName, msg.Message, msg.Title, msg.Album, msg.Artist, msg.User.Email, msg.User.Phone)
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