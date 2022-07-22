package router

import (
	"encoding/json"
	"github.com/pavel-one/TextRecognition/internal/responses"
	"net/http"
)

func responseJson(w http.ResponseWriter, r interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(r)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadGateway)

	response := responses.ErrorResponse{
		Errors: []string{err.Error()},
	}

	r, _ := json.Marshal(response)

	w.Write(r)
}
