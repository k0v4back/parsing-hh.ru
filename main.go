package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"./core/search"
	"./core/elastic"
	"./core"
)

type Response struct {
	Text 	string	`json:"text"`
	Error 	string	`json:"error"`
}

func Main(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r, "./static/forms/getAreaForm.html")
}

func Test(w http.ResponseWriter, r *http.Request)  {
		//area := r.FormValue("area")

		//result := search.InfoAboutOneCountry(area)
		//if result!= nil {
		//	//search.DataOfArea(result)
		//	fmt.Println(result.Url)
		//}
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


func main() {
	http.HandleFunc("/test", Test)
	http.HandleFunc("/", Main)
	http.HandleFunc("/upload/countries", core.JsonMiddleware(UploadCountries))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}