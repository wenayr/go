package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	buf := kitty{
		kitty2{
			name21:  "2dsd3",
			Name31:  "32dsd3",
			NOName1: "3dsd2"},
		"mmmmm",
		"Mya",
		"Mya3"}
	byt, _ := json.Marshal(buf)

	id := tId{"st"}
	order := tOrder2{
		tOrder{
			id,
			tMenuItem{
				id,
				345},
		},
		434,
		434}

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, string(byt)+"\n "+StructToString(order)+StructToString(byt))
	})
	http.ListenAndServe(port, nil)
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
