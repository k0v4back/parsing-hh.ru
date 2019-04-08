package search

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Countries struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Url 	string `json:"url"`
}

type ResponseCountry struct {
	Id 		string 	`json:"id"`
	Name 	string 	`json:"name"`
	Url 	string 	`json:"url"`
	Error 	string 	`json:"error"`
}

func IdOfAreaCountry(country string) ResponseCountry {
	resp, err := http.Get("https://api.hh.ru/areas/countries")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var c []Countries
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range c {
		if v.Name == country {
			return ResponseCountry{v.Id, v.Name, v.Url, nil}
		}
		return ResponseCountry{nil, nil, nil, "404"}
	}
	return ResponseCountry{nil, nil, nil, "404"}
}
