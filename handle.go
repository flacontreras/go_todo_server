package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func todos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var todos []Todo
	Db.Find(&todos)

	output, err := json.MarshalIndent(&todos, "", "\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, _ = w.Write(output)
}

func addTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	len := r.ContentLength
	body := make([]byte, len)
	_, _ = r.Body.Read(body)

	var todo Todo
	_ = json.Unmarshal(body, &todo)
	Db.Create(&todo)

	output, err := json.MarshalIndent(&todo, "", "\t")
	if err != nil {
		return
	}

	_, _ = w.Write(output)
}

func deleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return
	}

	var todo Todo
	Db.Where("id = ?", id).First(&todo)
	Db.Delete(&todo)
}
