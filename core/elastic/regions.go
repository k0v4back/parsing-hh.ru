package elastic

import (
	"context"
)

type Region struct {
	Id 			string 	`json:"id"`
	ParentId 	string 	`json:"parent_id"`
	Name 		string 	`json:"name"`
}

const mappingRegion = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"region":{
			"properties":{
				"id":{
					"type":"keyword"
				},
				"parent_id":{
					"type":"keyword"
				},
				"name":{
					"type":"text"
				}
			}
		}
	}
}`

func PutRegion(id, parentId, name string) error {
	ctx := context.Background()

	client := ConnectToElastic()

	checkRegion("regions")

	region := Region{Id: id, ParentId: parentId, Name: name}
	_, e := client.Index().
		Index("regions").
		Type("region").
		Id(id).
		BodyJson(region).
		Do(ctx)
	if e != nil {
		return e
	}
	return nil
}

func checkRegion(index string)  {
	ctx := context.Background()

	client := ConnectToElastic()

	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(index).BodyString(mappingRegion).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}