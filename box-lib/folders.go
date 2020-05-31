package box_lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type idRO struct {
	TotalCount int `json:"total_count"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	EntryType string `json:"type"`
	ID string `json:"id"`
	Etag string `json:"etag"`
}

//11056063660
func GetFileIDsToUpload(token string, folderID int, date string) ([]int, error) {
	idList := make([]int, 0)
	limit := 200
	offset := 0
	inProgress := true
	for inProgress {
		resp, err := getPageOfIds(token, folderID, date, limit, offset)
		if err != nil {
			return nil, err
		}
		for _, item := range resp.Entries {
			id, err := strconv.Atoi(item.ID)
			if err != nil {
				return nil, err
			}
			idList = append(idList, id)
		}
		offset+=limit
		inProgress = len(resp.Entries)==limit
	}
	fmt.Printf("found %v files", len(idList))
	return idList, nil
}

func getPageOfIds(token string, folderID int, date string, limit int, offset int) (idRO, error){
	client := &http.Client{
	}
	url := fmt.Sprintf("https://api.box.com/2.0/search?query=mp3&ancestor_folder_ids=%v&fields=id&offset=%v&limit=%v&created_at_range=%v,",folderID, offset, limit, date)
	fmt.Println(url)
	fmt.Println(token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return idRO{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := client.Do(req)
	if err != nil {
		return idRO{}, err
	}
	var RO idRO
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		err := json.NewDecoder(resp.Body).Decode(&RO)
		if err != nil {
			return idRO{}, err
		}
		return RO, nil
	}
	return RO, errors.New(fmt.Sprintf("GetSongIDs box api call returns %v", resp.Status))
}