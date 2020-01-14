package main

import (
	"fmt"
	"net/http"
)

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello World Yeah"))
}

func main() {
	port := "8080"

	http.HandleFunc("/hai", helloHandler)

	fmt.Println("Starting Web Server At Port : " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
