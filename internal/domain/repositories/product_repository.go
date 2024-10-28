package repositories

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/faruqii/goproto/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *entities.Product) error
	SearchProducts(query string) ([]entities.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
	es *elasticsearch.Client
}

func NewProductRepository(db *gorm.DB, es *elasticsearch.Client) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db, es: es}
}

func (r *ProductRepositoryImpl) CreateProduct(product *entities.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	// Index the product in Elasticsearch after creation
	r.indexProduct(product)
	return nil
}

func (r *ProductRepositoryImpl) indexProduct(product *entities.Product) {
	data, _ := json.Marshal(product)
	req := esapi.IndexRequest{
		Index:      "products",
		DocumentID: product.Id,
		Body:       strings.NewReader(string(data)), // Use strings.NewReader
		Refresh:    "true",
	}
	_, err := req.Do(context.Background(), r.es)
	if err != nil {
		// Handle error
	}
}

func (r *ProductRepositoryImpl) SearchProducts(query string) ([]entities.Product, error) {
	// Build the request body for the search
	body := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": query,
			},
		},
	}

	data, _ := json.Marshal(body) // Marshal the body to JSON
	res, err := r.es.Search(
		r.es.Search.WithContext(context.Background()),
		r.es.Search.WithIndex("products"),
		r.es.Search.WithBody(strings.NewReader(string(data))), // Use strings.NewReader
		r.es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var searchResults struct {
		Hits struct {
			Hits []struct {
				Source entities.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&searchResults); err != nil {
		return nil, err
	}

	var products []entities.Product
	for _, hit := range searchResults.Hits.Hits {
		products = append(products, hit.Source)
	}

	return products, nil
}
