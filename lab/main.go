package main

import (
	"fmt"
	"net/http"
)

var port = ":8000"

func main() {
	http.HandleFunc("/golab/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello")
	})
	http.ListenAndServe(port, nil)
}
