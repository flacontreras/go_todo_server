package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Id        int    `json:"id" sql:"index"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo
var idGenerator uuid.UUID

func handleOptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test cors")
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

	idGenerator = uuid.New()
	todos = []Todo{
		Todo{Id: 1, Title: "Todo One", Completed: false},
		Todo{Id: 2, Title: "Todo Two", Completed: false},
		Todo{Id: 3, Title: "Todo Three", Completed: false},
	}

	mux := httprouter.New()
	mux.GlobalOPTIONS = http.HandlerFunc(handleOptions)
	mux.GET("/todos", getTodos)
	mux.POST("/todos", addTodo)
	mux.DELETE("/todos/:id", deleteTodo)

	server := http.Server{
		Addr:    "127.0.0.1:7001",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
		return
	}
}
