package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var port = ":8000"

type any interface {
}
type fErr interface {
	f(er error)
}

func StructToString2(a any, fun fErr) string {
	buf, err := json.Marshal(a)
	if err == nil {
		return string(buf)
	}
	fun.f(err)
	return ""
}

func StructToString(a any) string {
	buf, _ := json.Marshal(a)
	return string(buf)
}

func main() {

	Router()
	http.HandleFunc("/hello", func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, "gggggggggggggg")
	})
	http.HandleFunc("/hello1", func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(writer, "111111111111")
	})

	http.ListenAndServe(port, nil)
}

func handleHelloWorld(write http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(write, "Hello!!!!!!")
}

type kitty2 struct {
	name21  string `json:"n1"`
	Name31  string
	NOName1 string `json:"name222"`
}
type kitty struct {
	kitty2
	name2  string `json:"name222"`
	Name3  string
	NOName string `json:"name"`
}

type tId struct {
	Id string `json:"id"`
}

type tMenuItem struct {
	tId
	Quantity int `json:"quantity"`
}

type tOrder struct {
	tId
	MenuItems tMenuItem `json:"menuItems"`
}

type tOrder2 struct {
	tOrder
	OrderedAtTimestamp int `json:"orderedAtTimestamp"`
	Cost               int `json:"cost"`
}

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("").Subrouter()
	s.HandleFunc("/hello3", handleHelloWorld).Methods(http.MethodGet)

	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beginTime := time.Now()
		h.ServeHTTP(w, r)

		endTime := time.Now()
		log.WithFields(log.Fields{
			"method":      r.Method,
			"url":         r.URL,
			"remoteAddr":  r.RemoteAddr,
			"userAgent":   r.UserAgent(),
			"elapsedTime": endTime.Sub(beginTime),
		}).Info("got a new request")
	})
}
