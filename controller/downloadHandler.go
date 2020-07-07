package controller

import (
	"fmt"
	box_lib "github.com/annakertesz/cp-music-lib/box-lib"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func download(db *sqlx.DB, token string, w http.ResponseWriter, r *http.Request) error{
	param:= chi.URLParam(r, "boxID")
	fmt.Println("download")
	fmt.Println(param)
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("couldnt convert string to int")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	resp, contentType, err := box_lib.DownloadFileBytes(token, id)
	if err != nil {
		fmt.Println("Couldnt download song")
		fmt.Println(err.Error())
		return err
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "attachment")
	w.Write(resp)
	return nil
}


