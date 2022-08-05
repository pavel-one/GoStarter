package router

import (
	"github.com/pavel-one/GoStarter/internal/responses"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	responseJson(w, responses.VersionResponse{Version: "1.0", Name: "Microservice for text recognition"})
}
