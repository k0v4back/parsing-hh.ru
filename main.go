package main

import (
	"net/http"
	"fmt"
	//"./core/search"
	"./core/elastic"
)

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

		//result := search.InfoAboutAllCountries()
		//for _, v := range *result {
		//	fmt.Println(v.Name)
		//}

		elastic.Test()

		//fmt.Println(result)
}

func main() {
	http.HandleFunc("/test", Test)
	http.HandleFunc("/", Main)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}