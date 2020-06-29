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
		w.Write([]byte("Only post supported"))
		return
	}
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read request. Error: " + err.Error()))
		return
	}
	response, err := api.PingPong(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read request. Error: " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/", httpResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
