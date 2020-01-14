package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(resp http.ResponseWriter, req *http.Request) {

	resp.Write([]byte("Method yang digunakan adalah : " + req.Method + ", "))

	resp.Write([]byte("Hello "))
	name := req.FormValue("namaUser")
	resp.Write([]byte(name))

	valGender := req.FormValue("gender")

	jnsKelamin := ""

	if strings.ToUpper(valGender) == "M" {
		jnsKelamin = "Laki - Laki"
	} else if strings.ToUpper(valGender) == "F" {
		jnsKelamin = "Perempuan"
	}
	resp.Write([]byte(", Jenis Kelamin Anda Adalah : " + jnsKelamin))

	// http://localhost:8080/hai?namaUser=gatra&gender=m Query url nya untuk di panggil di browser

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
