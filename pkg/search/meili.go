package search

import (
	"fmt"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

// Client is the global Meilisearch client instance
var Client *meilisearch.Client

// Init initializes the Meilisearch client using environment variables
func Init() {
	fmt.Println("Meilisearch Init called") // 👈 TEST
	Client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",       // Meilisearch instance URL
		APIKey: os.Getenv("MEILI_MASTER_KEY"), // Access key from .env file
	})
}
