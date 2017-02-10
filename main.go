package main

import (
	"github.com/chaselengel/checkin/mailer"
	"github.com/chaselengel/checkin/nodestore"
	"github.com/gorilla/handlers"
	"net/http"
)

var ns *nodestore.NodeStore
var mail *mailer.Mailer

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	ns, err = nodestore.Open()
	check(err)
	mail, err = mailer.Open()
	check(err)
	router := NewRouter()
	http.ListenAndServe(":8080", handlers.CORS()(router))
}
