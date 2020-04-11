package box_lib

import (
	"bytes"
	"fmt"
	"github.com/dhowden/tag"
	"io/ioutil"
	"testing"
)

func TestDownload(t *testing.T) {
	DownloadFileBytes("iO0Eh6etr9b1wnEy9Oksbcjfhyn8Wi0M", 133660155485)
}

func TestUploadFile(t *testing.T) {
	readFile, err := ioutil.ReadFile("../sources/instr.mp3")
	reader := bytes.NewReader(readFile)
	metadata, err := tag.ReadFrom(reader)
	picture := metadata.Picture()
	if err != nil {
		panic("panic")
	}
	file, _ := UploadFile("x6u5mPHFoGp2J1NNhJwa2ujkqN7khbAM", 110166546915, 8, picture.Data)
	fmt.Println(file)
}
