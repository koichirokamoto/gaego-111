package handlers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
)

// Fetch fetches web contents.
func Fetch(w http.ResponseWriter, r *http.Request) {
	api := "https://api.openweathermap.org/data/2.5/weather?q=Tokyo&units=metric&appid="
	apiKey := os.Getenv("WETHER_API_KEY")
	resp, err := http.Get(api + apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode >= 400 {
		http.Error(w, string(body), resp.StatusCode)
		return
	}

	main, _, _, err := jsonparser.Get(body, "main")
	if err != nil {
		http.Error(w, string(body), resp.StatusCode)
		return
	}

	w.Write(main)
}
