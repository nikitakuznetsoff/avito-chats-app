package main

import (
	"chatsapp/pkg/database"
	"chatsapp/pkg/handlers"

	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	dsn := "root:pass@tcp(db:3306)/chatsapp?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	repo := database.NewRepository(db)
	handler := handlers.Handler{ Repo: repo }

	router := mux.NewRouter()
	router.HandleFunc("/users/add", handler.CreateUser).Methods("POST")
	router.HandleFunc("/chats/add", handler.CreateChat).Methods("POST")
	router.HandleFunc("/chats/get", handler.GetChats).Methods("POST")
	router.HandleFunc("/messages/get", handler.GetMessages).Methods("POST")
	router.HandleFunc("/messages/add", handler.SendMessage).Methods("POST")

	address := ":9000"
	fmt.Printf("Starting server on port %v\n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}