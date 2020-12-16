package controller

import (
	"fmt"
	"github.com/annakertesz/cp-music-lib/services"
	box_lib "github.com/annakertesz/cp-music-lib/services/box-lib"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func download(db *sqlx.DB, token string, w http.ResponseWriter, r *http.Request) error{
	param:= chi.URLParam(r, "boxID")
	if len(param) < 5 {
		param = "736628393507"
	}
		fmt.Println("download")
	fmt.Println(param)
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("couldnt convert string to int")
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	resp, contentType, errModel := box_lib.DownloadFileBytes(token, id)
	if errModel != nil {
		services.HandleError(db, *errModel )
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "attachment")
	w.Write(resp)
	return nil
}


