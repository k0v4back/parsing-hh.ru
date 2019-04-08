package main

import (
	"net/http"
	"fmt"
)

func Main(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r, "./static/forms/getAreaForm.html")
}

func Test(w http.ResponseWriter, r *http.Request)  {
		area := r.FormValue("area")
		fmt.Println(area)
}

func main() {
	http.HandleFunc("/test", Test)
	http.HandleFunc("/", Main)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}