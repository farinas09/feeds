package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/farinas09/feeds-api/events"
	"github.com/farinas09/feeds-api/models"
	"github.com/farinas09/feeds-api/repository"
	"github.com/farinas09/feeds-api/search"
)

func onCreatedFeed(m events.CreatedFeedMessage) {
	feed := models.Feed{
		Id:          m.ID,
		Title:       m.Title,
		Description: m.Description,
	}
	if err := search.IndexFeed(context.Background(), feed); err != nil {
		log.Printf("Failed to index feed: %v", err)
	}
	log.Printf("Indexed feed: %s", feed.Id)
}

func ListFeedsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	feeds, err := repository.ListFeeds(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}

func SearchFeedsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query().Get("query")
	if len(query) == 0 {
		http.Error(w, "search query is required", http.StatusBadRequest)
		return
	}
	feeds, err := search.SearchFeeds(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}
