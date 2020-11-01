package box_lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/annakertesz/cp-music-lib/models"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"net/http"
)

func DownloadFileBytes(token string, id int) ([]byte, string, *models.ErrorModel) {
	responseBody, contentType, errModel := DownloadFile(token, id)
	if errModel != nil {
		return nil, "", errModel
	}
	defer responseBody.Close()
	bytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return nil, "", &models.ErrorModel{
			Service: "BoxLibService",
			Err:     err,
			Message: fmt.Sprintf("error while read download response"),
			Sev:     3,
		}
	}
	return bytes, contentType, nil

}

func DownloadFile(token string, id int) (io.ReadCloser, string, *models.ErrorModel) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.box.com/2.0/files/%v/content", id), nil)
	if err != nil {
		fmt.Println("error in download request")
		return nil, "", &models.ErrorModel{
			Service: "BoxLibService",
			Err:     err,
			Message: fmt.Sprintf("Error while creating request to download item: id = %v", id),
			Sev:     3,
		}
	}
	fmt.Println(req.URL)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	fmt.Println(req.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusOK {
		return resp.Body, resp.Header.Get("Content-Type"), nil
	}
	return nil, "",   &models.ErrorModel{
		Service: "BoxLibService",
		Err:     errors.New(fmt.Sprintf("error from downloader: %v", resp.Status)),
		Message: fmt.Sprintf("Error while download item: id = %v", id),
		Sev:     3,
	}
}

func UploadFile(token string, folderID int, filename int, file []byte) (int, error) {
	fmt.Println("upload cover photo")
	client := &http.Client{}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.SetBoundary("myBoundary")
	writer.WriteField("attributes", fmt.Sprintf("{\"name\":\"%v.jpg\", \"parent\":{\"id\":\"%v\"}}", filename, folderID))
	part, err := writer.CreateFormFile("file", "file.jpg")
	if err != nil {
		return 0, err
	}

	part.Write(file)

	err = writer.Close()
	if err != nil {
		return 0, err
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
	if resp.StatusCode == 201 {
	all, _ := ioutil.ReadAll(resp.Body)
	var result Result
	err = json.Unmarshal(all, &result)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	if len(result.Entries)<1 {
		return 0, errors.New("box doesnt return file id")
	}
	id, err := strconv.Atoi(result.Entries[0].Id)
	if err != nil {
		return 0, err
	}
	return id,nil}
	return 0, errors.New("file already exists")
}

type Result struct {
	Entries []struct{
		Id string `json:"id"`
	} `json:"entries"`
}


