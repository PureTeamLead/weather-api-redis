package main

import (
	"log"
	"net/http"
	"weather_api/pkg"
)

func main() {
	http.HandleFunc("/", pkg.WeatherHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
