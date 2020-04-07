package box_lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(token string, id int) error{
	client := &http.Client{
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.box.com/2.0/files/%v/content", id), nil)
	if err != nil {
		fmt.Println("error in download request")
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := client.Do(req)
	//fmt.Println(resp.Status)
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		// Create the file
		out, err := os.Create(fmt.Sprintf("../sources/music/%v.mp3", id))
		if err != nil {
			fmt.Println("error in oepning mp3 file in download")
			fmt.Println(err.Error())
			return errors.New("error")
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err!= nil {
			fmt.Println("error in writing mp3 file in download")

			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("error from downloader: %v %v", resp.Status, resp.Body))
}