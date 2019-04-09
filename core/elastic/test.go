package elastic

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Countries struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Url 	string `json:"url"`
}
const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"country":{
			"properties":{
				"id":{
					"type":"keyword"
				},
				"name":{
					"type":"keyword"
				},
				"url":{
					"type":"text",
					"store": true
				}
			}
		}
	}
}`

func Test()  {
	ctx := context.Background()

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	info, code, err := client.Ping("http://172.19.0.3:9300").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}