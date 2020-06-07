package controller

import (
	"encoding/json"
	"github.com/annakertesz/cp-music-lib/models"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func getAllTag(db *sqlx.DB, w http.ResponseWriter, r *http.Request){
	tags, err := models.GetTag(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tagJSON, err := json.Marshal(tagROListFromTag(tags))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(tagJSON)
	w.WriteHeader(http.StatusOK)
}

func tagROListFromTag(tags []models.Tag) []models.TagRO {
	tagROs := make([]models.TagRO, 0)
	for _, tag := range tags {
		tagROs = append(tagROs, models.TagRO{
			ID:   tag.TagID,
			Name: tag.TagName,
		})
	}
	return tagROs
}
