package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

const (
	POST_INDEX = "post"
	USER_INDEX = "user"
	ES_URL     = "http://10.182.0.2:9200"
)

func main() {
	// connection
	client, err := elastic.NewClient(elastic.SetURL(ES_URL))
	if err != nil {
		// Handle error
		panic(err)
	}
	defer client.Stop()

	// check if database exists
	exists, err := client.IndexExists(POST_INDEX).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	// index do not exist, create the index
	if !exists {
		fmt.Println("post index")
		mapping := `{
                        "mappings": {
				"properties": {
		                        "user":     { "type": "keyword", "index": false },
                                        "message":  { "type": "keyword", "index": false },
                                        "location": { "type": "geo_point" },
                                        "url":      { "type": "keyword", "index": false },
                                        "type":     { "type": "keyword", "index": false },
                                        "face":     { "type": "float" }
                                }
                        }
                }`
		_, err := client.CreateIndex(POST_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			fmt.Println("err2")
			panic(err)
		}
	}

	exists, err = client.IndexExists(USER_INDEX).Do(context.Background())
	if err != nil {
		fmt.Println("err")
		panic(err)
	}

	if !exists {
		fmt.Println("Debug")
		mapping := `{
                        "mappings": {
                                "properties": {
                                        "username": {"type": "keyword"},
                                        "password": {"type": "keyword", "index": false},
                                        "age":      {"type": "long", "index": false},
                                        "gender":   {"type": "keyword", "index": false}
                                }
                        }
                }`
		_, err := client.CreateIndex(USER_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Index is created.")

}
