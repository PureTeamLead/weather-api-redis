package pkg

import (
	"html/template"
	"log"
	"net/http"
	"weather_api/api"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	var res *api.Response

	tmpl, err := template.ParseFiles("/templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}

	if err = tmpl.Execute(w, *res); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

	address := r.URL.Query().Get("location")
	date := r.URL.Query().Get("date")

	res, err = api.GetForecast(address, []string{date, ""})
	if err != nil {
		log.Fatal(err)
	}
}
