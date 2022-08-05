package main

import (
	"fmt"
	"github.com/pavel-one/GoStarter/internal/queue"
	"github.com/pavel-one/GoStarter/internal/router"
	"log"
	"net/http"
)

func main() {
	q := queue.Run()

	route := new(router.Router)
	route.Worker = &q

	http.HandleFunc("/", route.Index)
	http.HandleFunc("/job", route.Test)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error start server: %s", err.Error()))
		return
	}
}
