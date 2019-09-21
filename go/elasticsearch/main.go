package main

import (
	"log"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func main() {
	endpoint := os.Getenv("ELASTICSEARCH_ENDPOPINT")
	username := os.Getenv("ELASTICSEARCH_USERNAME")
	password := os.Getenv("ELASTICSEARCH_PASSWORD")

	config := elasticsearch.Config{
		Addresses: []string{endpoint},
		Username:  username,
		Password:  password,
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("failed to es client: %+v\n", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Failed to get info %+v\n", err)
	}

	log.Println(res)
}
