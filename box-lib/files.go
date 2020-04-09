package box_lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DownloadFile(token string, id int) ([]byte, error) {
	client := &http.Client{
	}
	fmt.Println("download file")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.box.com/2.0/files/%v/content", id), nil)
	if err != nil {
		fmt.Println("error in download request")
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		//fmt.Println("create file")
		//file, err := os.Open(fmt.Sprintf("../sources/music/tmp.mp3"))
		//if err != nil {
		//	fmt.Println("error in creating mp3 file in download")
		//	fmt.Println(err.Error())
		//	return errors.New("error")
		//}
		//defer file.Close()
		//_, err = file.WriteAt([]byte{}, 0)
		//if err != nil {
		//	fmt.Println("error in cleaning mp3 file in download")
		//	fmt.Println(err.Error())
		//	return err
		//}
		//// Write the body to file
		//fmt.Println("write body to file")
		//bytes, err := ioutil.ReadAll(resp.Body)
		//_, err = file.WriteAt(bytes, 0)
		//if err != nil {
		//	fmt.Println("error in writing mp3 file in download")
		//	fmt.Println(err.Error())
		//	return err
		//}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("couldnt read bytes from response body")
		}
		return bytes, nil
	}
	return nil, errors.New(fmt.Sprintf("error from downloader: %v %v", resp.Status, resp.Body))
}
