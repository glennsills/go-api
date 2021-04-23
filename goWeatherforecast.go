package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func weatherForecastHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		forecasts := make([]*Weatherforecast, 5)
		for i := 0; i < 5; i++ {
			now := time.Now()
			temperatureC := rand.Intn(75) - 20
			temperatureF := int(float32(temperatureC)/float32(0.5556)) + 32
			forecasts[i] = &Weatherforecast{
				Date:         now.Format("Mon Jan 2 15:04:05 -0700 MST 2006"),
				TemperatureC: temperatureC,
				TemperatureF: temperatureF,
				Summary:      "Alright",
			}
		}
		forecastsJSON, _ := json.Marshal(forecasts)
		w.Write([]byte(forecastsJSON))
		return
	}

}

func main() {
	http.HandleFunc("/weatherForecast", weatherForecastHandler)
	http.ListenAndServe(":5000", nil)
}
