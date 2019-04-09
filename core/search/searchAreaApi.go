package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const NotFound  =  "Not found"

type Countries struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Url 	string `json:"url"`
}

type Area struct {
	Id string `json:"id"`
	ParentId string `json:"parent_id"`
	Name string `json:"name"`
	Areas []Region
}

type Region struct {
	Id string `json:"id"`
	ParentId string `json:"parent_id"`
	Name string `json:"name"`
	Areas []City
}

type City struct {
	Id string `json:"id"`
	ParentId string `json:"parent_id"`
	Name string `json:"name"`
}

func InfoAboutOneCountry(country string) *Countries {
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
			return &Countries{v.Id, v.Name, v.Url}
		}
		return nil
	}
	return nil
}


func InfoAboutAllCountries() *[]Countries {
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

	return &c
//	for _, v := range c {
//		if v.Name == country {
//			return &Countries{v.Id, v.Name, v.Url}
//		}
//		return nil
//	}
//	return nil
}


func DataOfArea(country_id string) {
	resp, err := http.Get("https://api.hh.ru/areas/" + country_id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var c Area
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	//for _, v := range c.Areas {
	//	fmt.Println(v.Name)
	//}

	for _, v := range c.Areas {
		regions := v.Areas
		for _, d := range regions {
			fmt.Printf("City %s; id - %s ", d.Name, d.Id)
		}
	}

	//for _, v := range c {
	//	if v.Name == country {
	//		return ResponseCountry{v.Id, v.Name, v.Url, nil}
	//	}
	//	return ResponseCountry{nil, nil, nil, "404"}
	//}
	//return ResponseCountry{nil, nil, nil, "404"}
}