package models

import (
	"fmt"
	"testing"
)

func TestSessionVAlidation(t *testing.T){
	db, _ := connectToDB()
	id, err := ValidateSessionID(db, "9ca062f0-23db-40c4-8226-40aa5fc4c7b3")
	fmt.Println(id, err)
}
