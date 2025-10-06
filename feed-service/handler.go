package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/farinas09/feeds-api/events"
	"github.com/farinas09/feeds-api/models"
	"github.com/farinas09/feeds-api/repository"
	"github.com/segmentio/ksuid"
)

type createdFeedRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func createFeedHandler(w http.ResponseWriter, r *http.Request) {
	var req createdFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdAt := time.Now()
	id, err := ksuid.NewRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	feed := models.Feed{
		Id:          id.String(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   createdAt,
	}
	err = repository.InsertFeed(r.Context(), &feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := events.PublishCreatedFeed(r.Context(), &feed); err != nil {
		log.Printf("Failed to publish created feed: %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&feed)
}
