package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ScriptIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ns.Nodes)
}

func ScriptCheckin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["scriptName"]
	node, err := ns.InsertOrUpdate(name)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(node)
}

func ScriptDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["scriptId"])
	err := ns.Delete(id)
	result := true
	if err != nil {
		result = false
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ScriptShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["scriptId"])
	_, node := ns.FindById(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(node)
}

func ScriptSchedule(w http.ResponseWriter, r *http.Request) {
}

func MailIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mail.Emails)
}

func MailCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	result := true
	if mail.Insert(address) != nil {
		result = false
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func MailDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["mailId"])
	err := mail.Delete(id)
	result := true
	if err != nil {
		result = false
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
