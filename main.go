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
		cash := struct {
			Message string
		} {
			"This was a GET, cuz",
		}
		RespondWithJSON(w, http.StatusOK, cash)
	case "POST":
		type Request struct {
			Guess string `json:"guess"`
		}
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println("Error decoding request: ", err.Error())
			RespondWithError(w, http.StatusInternalServerError, "This request was not that tight.  I think it was your fault, not mine. Don't @ me")
			return
		}
		defer r.Body.Close()

		cash := struct {
			Message string
		} {
			"UHHH...YUP!!!!",
		}
		if req.Guess == "Tom Brady" {
			RespondWithJSON(w, 200, cash)
		} else {
			RespondWithJSON(w, 200, map[string]string{"Message": "LOL NAH"})
		}
	}
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