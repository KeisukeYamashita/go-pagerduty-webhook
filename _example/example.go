package main

import (
	"encoding/json"
	"net/http"

	webhook "github.com/KeisukeYamashita/go-pagerduty-webhook"
)

func main() {
	http.HandleFunc("/", incidentHandler())
	http.ListenAndServe(":8080", nil)
}

func incidentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var payload webhook.Payload
		err := decoder.Decode(&payload)
		if err != nil {
			return
		}
	}
}
