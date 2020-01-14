package main

import (
	"fmt"
	mdw "materiGolang/goAPI/BelajarModular/middleware"
	adv "materiGolang/goAPI/BelajarModular/middleware/advance"
	"net/http"
)

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(3)
	resp.Write([]byte("Hello World"))
}

func main() {
	port := "8080"

	// http.Handle("/handler", mdw.Logger(http.HandlerFunc(helloHandler))) // memanggil middleware modular
	// http.HandleFunc("/handler2", adv.Logger2(helloHandler))             // memanggil advance modular
	http.Handle("/handler3", mdw.Logger(adv.Logger2(http.HandlerFunc(helloHandler))))
	fmt.Println("Starting Web Server At Port : " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
