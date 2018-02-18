package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/iamgreggarcia/gofigure/api/config"
	"github.com/iamgreggarcia/gofigure/api/model"
)

// Pointer to our config to be passes from Main
var cfg *config.Configuration

// Initialize sets the handler config environment
func Initialize(config *config.Configuration) {
	cfg = config
}

// GetHelloHandler is a handler that returns JSON:
// { "message": "Hello, <name>!" }
func GetHelloHandler(w http.ResponseWriter, r *http.Request) {
	// We use the github.com/fatih/color package to
	// colorize our logs.
	// Use mux to parse request for the "name" parameter
	// and use it to create a Model.Message struct for use
	// in our response
	vars := mux.Vars(r)
	name := vars["name"]
	message := fmt.Sprintf("Hello, %s!", name)
	m := model.Message{Message: message}

	color.Set(color.FgCyan)

	// Log results
	log.Printf("Request received: /api/v1/greetings/hello/%s \n", name)

	color.Unset()
	color.Set(color.FgCyan)

	log.Printf(`Response: {"message":"%s"}`, message)

	color.Unset()
	color.Set(color.FgHiGreen)

	log.Println("Sever listening on port:", cfg.Server.Port)

	color.Unset()

	// Return response in JSON
	respondWithJSON(w, http.StatusOK, m)
}

// RespondWithError returns JSON encoded error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})

}

// RespondWithJSON returns JSON for error handling
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
