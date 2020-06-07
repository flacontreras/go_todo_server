package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Id        int    `json:"id" sql:"index"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		// Set CORS headers
		header := w.Header()

		header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*")
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := httprouter.New()
	mux.GlobalOPTIONS = http.HandlerFunc(handleOptions)
	mux.GET("/todos", todos)
	mux.POST("/todos", addTodo)
	mux.DELETE("/todos/:id", deleteTodo)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
		return
	}
}
