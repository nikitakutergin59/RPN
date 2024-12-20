package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nikitakutrergin59/RPN/pkg/calculation/calc"
)

type Request struct {
	Expression string `json:"expression" `
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func cakculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request bodt", http.StatusBadRequest)
	}

}
