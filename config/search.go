// pkg/config/search.go

package config

import (
	"context"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

var ESClient *elastic.Client

func InitElasticsearch() {
	var err error
	ESClient, err = elastic.NewClient(
		elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error initializing Elasticsearch client: %v", err)
	}
}

func CreateIndex(index string, mapping string) error {
	ctx := context.Background()
	exists, err := ESClient.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		_, err = ESClient.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func IndexDocument(index string, id string, doc interface{}) error {
	ctx := context.Background()
	_, err := ESClient.Index().
		Index(index).
		Id(id).
		BodyJson(doc).
		Do(ctx)
	return err
}
