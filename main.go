package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func main() {
	http.Handle("/foo", fooHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type fooHandler struct{}

func (f fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := struct {
		Animal string
		Genus string
		Species string
		Type string
	} {
		"Goat", 
		"Thomas", 
		"Brady", 
		"Rare AF",
	}
	response, _ := json.Marshal(cash)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(response)
}