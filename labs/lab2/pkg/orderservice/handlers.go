package orderservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type any interface{}

func StructToStringAbs(a any) string {
	buf, _ := json.Marshal(a)
	//if (err!=nil) {http.Error(w, err.Error(), http.StatusInternalServerError)}// по факту надо принтить трасировку
	return string(buf)
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
	MenuItems []tMenuItem `json:"menuItems"`
}

type tOrder2 struct {
	tOrder
	OrderedAtTimestamp int `json:"orderedAtTimestamp"`
	Cost               int `json:"cost"`
}

type kitty struct {
	Name string `json:"name"`
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello-world", handleHelloWorld).Methods(http.MethodGet)
	r.HandleFunc("/cat", handleKitty).Methods(http.MethodGet)

	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/order/{ID}", handleOrder).Methods(http.MethodGet)
	s.HandleFunc("/orders", handleOrders).Methods(http.MethodGet)

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

func handleKitty2(w http.ResponseWriter, request *http.Request) {
	cat := kitty{"Кот"}
	b, _ := json.Marshal(cat)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func handleKitty(w http.ResponseWriter, _ *http.Request) {
	cat := kitty{"Кот"}
	mByt, err := json.Marshal(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(mByt)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func handleHelloWorld(write http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(write, "Hello!")
}

func handleOrder(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	var st string = vars["ID"] // ???
	id := tId{st}
	menu := []tMenuItem{{id, 45}}
	//order :=tOrder{id,menu}

	data := tOrder2{
		tOrder{id, menu},
		434,
		434}
	result, _ := json.Marshal(data)

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(writer, string(result)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}

func handleOrders(w http.ResponseWriter, _ *http.Request) {

	id := tId{"23ds23df3f"}
	menuItems := []tMenuItem{{id, 1}}
	orders := []tOrder{{id, menuItems}}

	result, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(result)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
