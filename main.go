package main

import (
	"fmt"
	"github.com/chaselengel/checkin/mailer"
	"github.com/chaselengel/checkin/nodestore"
	"net/http"
)

var ns *nodestore.NodeStore
var mail *mailer.Mailer

func main() {
	var err error
	ns, err = nodestore.Open()
	if err != nil {
		fmt.Println("Failed to open node store: " + err.Error())
		return
	}
	mail, err = mailer.Open()
	if err != nil {
		fmt.Println("Failed to open mail store: " + err.Error())
		return
	}
	router := NewRouter()
	http.ListenAndServe(":8080", router)
}
