package pkg

import (
	"html/template"
	"log"
	"net/http"
	"weather_api/api"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		http.Error(w, "Error parsing files", http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	address := r.FormValue("location")
	date := r.FormValue("date")

	parseObj, err := api.GetForecast(address, []string{date, ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if parseObj.Description == "" {
		http.Error(w, "Sorry, but 3rd party API cannot fetch info about this location", http.StatusBadRequest)
		return
	}

	log.Println(parseObj.CurrentConditions.Temp)

	tmpl, err := template.ParseFiles("../templates/results.html")
	if err != nil {
		http.Error(w, "Error parsing files", http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, parseObj); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
