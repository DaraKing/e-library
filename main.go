package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/images/{rest}", http.StripPrefix("/images/", http.FileServer(http.Dir( "./images"))))

	myRouter.HandleFunc("/get-reserved-books", getReservedBooks).Methods("GET")
	myRouter.HandleFunc("/get-reserved-books", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/login", loginFunc).Methods("POST")
	myRouter.HandleFunc("/login", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/register", registerFunc).Methods("POST")
	myRouter.HandleFunc("/register", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/userInfo", getUserInfo).Methods("POST")
	myRouter.HandleFunc("/userInfo", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/add-book", addBook).Methods("POST")
	myRouter.HandleFunc("/add-book", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/get-books", getBooks).Methods("GET")
	myRouter.HandleFunc("/get-books", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/book/{id}", getSingleBook).Methods("GET")
	myRouter.HandleFunc("/book/{id}", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/rate", bookRating).Methods("POST")
	myRouter.HandleFunc("/rate", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/take-book", takeBook).Methods("POST")
	myRouter.HandleFunc("/take-book", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/send-comment", getComment).Methods("POST")
	myRouter.HandleFunc("/send-comment", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/set-as-return/{id}", setAsReturn).Methods("DELETE")
	myRouter.HandleFunc("/set-as-return/{id}", optionsRequest).Methods("OPTIONS")

	log.Fatal(http.ListenAndServe(":3030", myRouter))
}

func main() {
	fmt.Println("Dara API")
	handleRequests()
}