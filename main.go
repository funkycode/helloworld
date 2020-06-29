package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/funkycode/helloworld/api"
)

func httpResponse(w http.ResponseWriter, r *http.Request) {
	var request []byte
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Only post supported")); err != nil {
			log.Fatalf("Failed to return error msg. Error: %s", err)
		}
		return
	}
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Failed to read request. Error: " + err.Error())); err != nil {
			log.Fatalf("Failed to return error msg. Error: %s", err)
		}
		return
	}
	response, err := api.PingPong(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Failed to read request. Error: " + err.Error())); err != nil {
			log.Fatalf("Failed to return error msg. Error: %s", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		log.Fatalf("Failed to return response. Error: %s", err)
	}
}

func main() {
	http.HandleFunc("/", httpResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
