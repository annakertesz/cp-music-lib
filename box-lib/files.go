package box_lib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"


	"net/http"
)

func DownloadFileBytes(token string, id int) ([]byte, string, error) {
	responseBody, contentType, err := DownloadFile(token, id)
	if err != nil {
		return nil, "",  err
	}
	defer responseBody.Close()
	bytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		fmt.Println("couldnt read bytes from response body")
		errors.New("couldnt read bytes from response body")
	}
	return bytes, contentType, nil

}

func DownloadFile(token string, id int) (io.ReadCloser, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.box.com/2.0/files/%v/content", id), nil)
	if err != nil {
		fmt.Println("error in download request")
		return nil, "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusOK {
		return resp.Body, resp.Header.Get("Content-Type"), nil
	}
	return nil, "",  errors.New(fmt.Sprintf("error from downloader: %v %v", resp.Status, resp.Body))
}

func UploadFile(token string, folderID int, image []byte) error {
	fmt.Println("upload cover photo")
	client := &http.Client{}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.SetBoundary("myBoundary")
	writer.WriteField("attributes", "{\"name\":\"Photo.jpg\", \"parent\":{\"id\":\"110166546915\"}}")
	part, err := writer.CreateFormFile("file", "file.jpg")
	if err != nil {
		return err
	}

	part.Write(image)

	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://upload.box.com/api/2.0/files/content"), body)
	if err != nil {
		fmt.Println("error in upload request")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Add("Content-Type", "multipart/form-data; boundary=myBoundary")
	fmt.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(all))
	return nil
}


