package controller

import (
	"github.com/annakertesz/cp-music-lib/models"
	"testing"
)

func TestEmailSender(t *testing.T){
	user := models.UserReqObj{
		Username:     "username",
		FirstName:    "First",
		LastName:     "Name",
		Email:        "email@email.com",
		Phone:        "345039483045",
	}
	sendVerifyEmail(user, "http://google.com")
}
