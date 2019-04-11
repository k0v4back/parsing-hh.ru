package elastic

import (
	"github.com/olivere/elastic"
	"log"
	"time"
	"context"
	"fmt"
)

type countries struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Url 	string `json:"url"`
}
const mapp = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"maps":{
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
	info, code, err := client.Ping("http://elasticsearch:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	//Set one country
	country1 := Countries{Id: "1", Name: "Russia", Url: "http://russia.com"}
	put1, err := client.Index().
		Index("countries").
		Type("country").
		Id("1").
		BodyJson(country1).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed country %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	//Get country
	get1, err := client.Get().
		Index("countries").
		Type("country").
		Id("1").
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	//Search
	termQuery := elastic.NewTermQuery("name", "Russia")
	searchResult, err := client.Search().
		Index("countries").   // search in index "countries"
		Query(termQuery).   // specify the query
		Sort("id", true). // sort by "id" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)


}