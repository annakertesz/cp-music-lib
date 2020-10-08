package box_lib

import (
	"fmt"
	"testing"
)

func TestFolder(t *testing.T) {
	upload, err := GetFileIDsToUpload("U5lXslsTv5YPAz3H2vJe6h824uS25tJi", 11056063660, "2014-02-01T00:00:00Z")
	if err != nil {
		fmt.Println(err)
	}
	for i := range upload {
	fmt.Println(upload[i])}
}
