package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	MIDDLEWARE
*/

// cara pertama
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Printf("URL %v dipanggil pada jam %v", req.URL.Path, time.Now())
		fmt.Println(1)
		next.ServeHTTP(resp, req) // akan menjalan fungsi yang memanggilnya, setelah selesai akan melanjutkan di bawah nya
		fmt.Println(2)
		fmt.Printf("URL %v dipanggil pada jam %v", req.URL.Path, time.Now())
	})
}

// cara kedua
func Logger2(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Header-x", "XXXXXXXX")
		fmt.Println(1)
		next.ServeHTTP(resp, req) // akan menjalan fungsi yang memanggilnya, setelah selesai akan melanjutkan di bawah nya
		fmt.Println(2)
		resp.Write([]byte("Done Testing"))
	})
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(3)
	resp.Write([]byte("Hello World"))
}

func main() {
	port := "8080"

	// untuk mendaftar kan ke url
	http.Handle("/hai", Logger(http.HandlerFunc(helloHandler))) // cara memanggil middlwate pertama
	http.HandleFunc("/haii", Logger2(helloHandler))             // cara memanggil middlwate Kedua

	fmt.Println("Starting Web Server At Port : " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
