package smtpClient

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/config"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailSender struct {
	smtpConfig config.SmtpEmailConfig
}

func NewEmailSender(config config.SmtpEmailConfig) EmailSender {
	return EmailSender{smtpConfig:	config}
}

func (sender *EmailSender) SendVerifyEmail(user models.UserReqObj, verifyEndpoint string) error {
	e := email.NewEmail()
	e.From = sender.smtpConfig.FromEmail
	e.To = []string{fmt.Sprintf("Central admin <%s>", sender.smtpConfig.ToEmail)}
	e.Subject = fmt.Sprintf("%v %v would like have access", user.FirstName, user.LastName)
	e.Text = []byte(fmt.Sprintf(
		"New user request: \n"+
			"Name: %v %v \n"+
			"Email: %v \n"+
			"Phone number: %v \n"+
			"\n"+
			"Please go to http://%v  to accept his/her request",
		user.FirstName, user.LastName, user.Email, user.Phone, verifyEndpoint))
	e.HTML = []byte(fmt.Sprintf(
		"<h2 style=\"color: #2e6c80;\">New user request:</h2>"+
			"<ul>"+
			"<li>Name: %v %v</li>"+
			"<li>Email: %v</li>"+
			"<li>Phone number: %v</li>"+
			"</ul>"+
			"<p>Please click <a href=\"http://%v\">here</a> to accept his/her request</p>",
		user.FirstName, user.LastName, user.Email, user.Phone, verifyEndpoint))
	return sender.sendEmail(e)
}

func (sender *EmailSender) SendSongBuyEmail(msg models.BuySongObj) error {
	e := email.NewEmail()
	e.From = sender.smtpConfig.FromEmail
	e.To = []string{fmt.Sprintf("Central admin <%s>", sender.smtpConfig.ToEmail)}
	e.Subject = fmt.Sprintf("%v %v would like to buy a song", msg.User.FirstName, msg.User.LastName)
	e.Text = []byte(fmt.Sprintf(
		"New song request:\n"+
			"Dear admin\n"+
			"%v %v would like to buy a song:\n"+
			"%v\n"+
			"Title: %v\n"+
			"Album: %v\n"+
			"Artist: %v\n"+
			"You can contact him/her via email (%v) or phone (%v)",
		msg.User.FirstName, msg.User.LastName, msg.Message, msg.Title, msg.Album, msg.Artist, msg.User.Email, msg.User.Phone))
	e.HTML = []byte(fmt.Sprintf(
		"<h2 style=\"color: #2e6c80;\">New song request:</h2>"+
			"<p>Dear admin</p>"+
			"<p>%v %v would like to buy a song:</p>"+
			"<p>%v</p>"+
			"<ul>"+
			"<li>Title: %v</li>"+
			"<li>Album: %v</li>"+
			"<li>Artist: %v</li>"+
			"</ul>"+
			"<p>You can contact him/her via email (%v) or phone (%v)</p>",
		msg.User.FirstName, msg.User.LastName, msg.Message, msg.Title, msg.Album, msg.Artist, msg.User.Email, msg.User.Phone))
	return sender.sendEmail(e)
}

func (sender *EmailSender) sendEmail(email *email.Email) error {
	err := email.Send(sender.smtpConfig.ServerAddress,
		smtp.PlainAuth("", sender.smtpConfig.UserName, sender.smtpConfig.Password, sender.smtpConfig.Host))
	if err != nil {
		panic(err)
	}
	return nil
}
