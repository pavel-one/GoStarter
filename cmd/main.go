package main

import (
	"github.com/pavel-one/TextRecognition/internal/router"
	"net/http"
)

func main() {
	http.HandleFunc("/", router.Index)

	http.ListenAndServe(":3000", nil)
}
