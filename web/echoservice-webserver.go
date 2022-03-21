package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"go.kimpton.io/echo/common"
)

type EchoWebServer struct{}

func main() {
	flag.Parse()

	log.Printf("Starting Echo Web Server on port %d", common.Port())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World - From Go!")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", common.Port()), nil))

}
