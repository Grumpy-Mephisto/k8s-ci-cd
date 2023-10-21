package config

import (
	"os-container-project/pkg/utils"

	"github.com/gofiber/fiber/v2/log"
	"github.com/olivere/elastic/v7"
)

type ElasticsearchConfig struct {
	URL       string
	IndexName string
	Username  string
	Password  string
}

func NewElasicSearchConfig() *ElasticsearchConfig {
	return &ElasticsearchConfig{
		URL:       utils.GetEnv("ELASTICSEARCH_URL", "http://localhost:9200").(string),
		IndexName: utils.GetEnv("ELASTICSEARCH_INDEX_NAME", "os-container-project").(string),
		Username:  utils.GetEnv("ELASTICSEARCH_USERNAME", "elastic").(string),
		Password:  utils.GetEnv("ELASTICSEARCH_PASSWORD", "changeme").(string),
	}
}

func (e *ElasticsearchConfig) NewElasticsearchClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(e.URL),
		elastic.SetBasicAuth(e.Username, e.Password),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error while creating elasticsearch client: %s", err.Error())
		return nil, err
	}

	return client, nil
}
