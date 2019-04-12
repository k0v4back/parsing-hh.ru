package elastic

import (
	"encoding/json"
	"github.com/olivere/elastic"
	"log"
	"time"
	"context"
	"fmt"
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

func checkCountries(index string)  {
	ctx := context.Background()

	client := ConnectToElastic()

	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func ConnectToElastic() *elastic.Client {
	ctx := context.Background()

	checkCountries("properties")


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

	client := ConnectToElastic()

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

func GetCountryByName(name string) *elastic.SearchResult {
	client := ConnectToElastic()

	ctx := context.Background()

	termQuery := elastic.NewTermQuery("name", name)
	searchResult, err := client.Search().
		Index("countries").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("name", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println(searchResult.Hits.TotalHits)

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Countries
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.Name, t.Id)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}



	return searchResult
}