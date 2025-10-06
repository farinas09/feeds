package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farinas09/feeds-api/database"
	"github.com/farinas09/feeds-api/events"
	"github.com/farinas09/feeds-api/repository"
	"github.com/farinas09/feeds-api/search"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	NatsAddress      string `envconfig:"NATS_ADDRESS" required:"true"`
	ElasticSearchURL string `envconfig:"ELASTICSEARCH_URL" required:"true"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to process config: %v", err)
	}

	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	repo, err := database.NewPostgresRepository(addr)
	if err != nil {
		log.Fatalf("Failed to create postgres repository: %v", err)
	}
	repository.SetRepository(repo)

	es, err := search.NewElasticSearchRepository(fmt.Sprintf("http://%s", cfg.ElasticSearchURL))
	if err != nil {
		log.Fatalf("Failed to create elasticsearch repository: %v", err)
	}
	search.SetSearchRepository(es)

	defer search.Close()

	n, err := events.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		log.Fatalf("Failed to create nats event store: %v", err)
	}
	events.SetEventStore(n)

	err = n.OnCreatedFeed(onCreatedFeed)
	if err != nil {
		log.Fatalf("Failed to subscribe to created feed: %v", err)
	}

	defer events.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/feeds", ListFeedsHandler).Methods(http.MethodGet)
	router.HandleFunc("/search", SearchFeedsHandler).Methods(http.MethodGet)
	return
}
