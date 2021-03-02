package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	orderservice "labs/lab2/pkg/orderservice"
	"net/http"
)

var port = ":8000"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	a:=orderservice.Handle()
	http.HandleFunc("/golab/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, a)
	})
	http.ListenAndServe(port, nil)
}
