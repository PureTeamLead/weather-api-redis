// handle errors on localhost more beautifully
package main

import (
	"log"
	"net/http"
	"weather_api/pkg"
)

func main() {
	http.HandleFunc("/", pkg.WeatherHandler)
	http.HandleFunc("/results", pkg.ResultsHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
