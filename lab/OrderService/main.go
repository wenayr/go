package main

import (
	"fmt"
	"net/http"
	"transfer/Router"
)

var port = ":8000"

func main() {
	r := Router()
	http.HandleFunc("/golab/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello")
	})
	http.ListenAndServe(port, nil)
}
