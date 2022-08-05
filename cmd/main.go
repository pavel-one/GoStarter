package main

import (
	"fmt"
	"github.com/pavel-one/GoStarter/internal/router"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", router.Index)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error start server: %s", err.Error()))
		return
	}
}
