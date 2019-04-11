package elastic

import (
	"github.com/olivere/elastic"
	"log"
	"time"
	"context"
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

func connectToElastic() *elastic.Client {
	ctx := context.Background()

	//Connect to elastic
	client, err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		time.Sleep(3 * time.Second)
	}

	//Ping elastic
	_, _, e := client.Ping("http://elasticsearch:9200").Do(ctx)
	if e != nil {
		panic(e)
	}

	return client
}

func PutCountry(id, name, url string) error {
	ctx := context.Background()

	client := connectToElastic()

	country := Countries{Id: id, Name: name, Url: url}
	_, e := client.Index().
		Index("countries").
		Type("country").
		Id(id).
		BodyJson(country).
		Do(ctx)
	if e != nil {
		return e
	}
	return nil
}