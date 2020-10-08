package services

import (
	"fmt"
	"log"
	"testing"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func TestEmailSender(t *testing.T) {
	SENDGRID_API_KEY:="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"

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

