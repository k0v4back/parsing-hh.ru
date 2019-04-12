package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"./core/search"
	"./core/elastic"
	"./core"
	//"reflect"
)

type Response struct {
	Text 	string	`json:"text"`
	Error 	string	`json:"error"`
}

func Main(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r, "./static/forms/getAreaForm.html")
}

func Test(w http.ResponseWriter, r *http.Request)  {
	elastic.GetCountryByName("Казахстан")
}

func UploadCountries(w http.ResponseWriter, r *http.Request)  {
	result := search.InfoAboutAllCountries()
	for _, v := range *result {
		err := elastic.PutCountry(v.Id, v.Name, v.Url)
		if err != nil {
			log.Fatal("Error with uploading countries to elastic")
		}
	}
	resp := &Response{Text:"All countries successfully loaded into Elastic", Error:"nil"}
	json.NewEncoder(w).Encode(resp)
}

func UploadRegions(w http.ResponseWriter, r *http.Request) {
	result := search.InfoAboutAllCountries()
	for _, v := range *result {
		result := search.DataOfAllRegionsInCountry(v.Id)
		for _, v := range *result {
			go func (){
				err := elastic.PutRegion(v.Id, v.ParentId, v.Name)
				if err != nil {
					log.Fatal("Error with uploading regions to elastic")
				}
			}()
		}
	}
}


func main() {
	http.HandleFunc("/test", Test)
	http.HandleFunc("/", Main)
	http.HandleFunc("/upload/countries", core.JsonMiddleware(UploadCountries))
	http.HandleFunc("/upload/regions", core.JsonMiddleware(UploadRegions))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}