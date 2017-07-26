package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// CityHandler The CityHandler
func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities': 'San Franciso, Amsterdam, Berlin'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)

}

func main() {
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
