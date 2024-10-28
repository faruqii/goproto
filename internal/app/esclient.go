package app

import (
	"fmt"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

func GetESClients() (*elastic.Client, error) {
	client, err := elastic.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	fmt.Println("ES client created")
	return client, nil
}
