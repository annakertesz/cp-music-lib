package box_lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(token string, id int) error {
	client := &http.Client{
	}
	fmt.Println("download file")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.box.com/2.0/files/%v/content", id), nil)
	if err != nil {
		fmt.Println("error in download request")
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		fmt.Println("create file")
		file, err := os.Create(fmt.Sprintf("sources/music/%v.mp3", id))
		if err != nil {
			fmt.Println("error in creating mp3 file in download")
			fmt.Println(err.Error())
			return errors.New("error")
		}
		defer file.Close()

		// Write the body to file
		fmt.Println("write body to file")
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println("error in writing mp3 file in download")
			fmt.Println(err.Error())
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("error from downloader: %v %v", resp.Status, resp.Body))
}
