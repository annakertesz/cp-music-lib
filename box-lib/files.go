package box_lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(token string, id int){
	client := &http.Client{
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.box.com/2.0/files/%v/content", id), nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := client.Do(req)
	//fmt.Println(resp.Status)
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		// Create the file
		out, err := os.Create(fmt.Sprintf("../sources/music/%v.mp3", id))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
	}
}