package main

import (
	"fmt"
	"github.com/chaselengel/checkin/datastore"
	"html/template"
	"net/http"
)

var ds, _ = datastore.Open()

// Insert/Update node timestamp.
func checkinHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/checkin/"):]

	ds.InsertOrUpdate(name)
}

// Delete node and timestamps from datastore.
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/delete/"):]

	ds.Delete(name)

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, ds.Nodes)
}

// List all timestamps for node's name.
func listHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/list/"):]

	_, node := ds.Find(name)
	if node == nil {
		fmt.Fprintf(w, "Failed to find "+name)
		return
	}

	t, _ := template.ParseFiles("list.html")
	t.Execute(w, node)
}

func scheduleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/schedule/"):]
	r.ParseForm()
	ds.SetSchedule(name, r.Form["digit"][0]+r.Form["schedule"][0])
}

// Show all nodes and newest timestamp.
func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, ds.Nodes)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/checkin/", checkinHandler)
	http.HandleFunc("/list/", listHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/schedule/", scheduleHandler)
	http.ListenAndServe(":8080", nil)
}
