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
	switch r.Method {
	case "GET":
		GetHandler(w, r)
	case "POST":
		PostHandler(w, r)
	}
}

type Response struct {
	Message string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{"This was a GET, ya dig?"}
	RespondWithJSON(w, http.StatusOK, res.Message)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Guess string `json:"guess"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding request: ", err.Error())
		RespondWithError(w, http.StatusInternalServerError, "This request was not that tight.  I think it was your fault, not mine. Don't @ me")
		return
	}
	defer r.Body.Close()
	
	res := CheckGuess(req.Guess)
	RespondWithJSON(w, http.StatusOK, Response{res})	
}

func CheckGuess(guess string) (answer string) {
	if guess == "Tom Brady" {
		answer = "UHHH...YUP!!!!"
	} else {
		answer = "LOL NAH"
	}
	return
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}