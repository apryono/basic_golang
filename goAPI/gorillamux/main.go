package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Get"))
	fmt.Println("Hello World")
}

func helloPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Post"))
	fmt.Println("Hello World")
}

func main() {
	port := "8080"

	router := mux.NewRouter()
	// cara pertama penggunaan method
	// router.HandleFunc("/hello", helloGetHandler).Methods("GET")
	// router.HandleFunc("/hello", helloPostHandler).Methods("POST")

	//cara method kedua
	router.HandleFunc("/hello", helloGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/hello", helloPostHandler).Methods(http.MethodPost)

	fmt.Println("Starting Web Server At Port : " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Opps")
		fmt.Println(err.Error())
	}
}
