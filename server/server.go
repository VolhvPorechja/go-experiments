package server

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
)

func Server() {
	s := serverConfig{
		name: "fuckyou",
		Address: Address{
			protocol: "http",
			address:  "localhost:8088",
			resource: "/",
		},
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(s.address, nil))
}
func handler(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("server/templates/main.html")
	if err != nil {
		serr := fmt.Sprintf("ERROR: %s", err)
		fmt.Fprintf(writer, serr)
		return
	}
	t.Execute(writer, request.URL.Path)
}

type Address struct {
	protocol, address, resource string
}

type serverConfig struct {
	Address
	name string
}

type tree struct {
	value       interface{}
	left, right *tree
}
