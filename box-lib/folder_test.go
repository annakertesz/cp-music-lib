package box_lib

import (
	"fmt"
	"testing"
)

func TestFolder(t *testing.T) {
	upload, err := GetFileIDsToUpload("AP8rFgRAD58dQxys4lmN8BVf8ffZUZW7", 11056063660, "2017-05-15T13:35:01-07:00")
	if err != nil {
		fmt.Println(err)
	}
	for i := range upload {
	fmt.Println(upload[i])}
}
