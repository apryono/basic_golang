package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func helloHandlerGet(resp http.ResponseWriter, req *http.Request) {

	name := req.FormValue("namaUser")

	valGender := req.FormValue("gender")

	jnsKelamin := ""

	if strings.ToUpper(valGender) == "M" {
		jnsKelamin = "Laki - Laki"
	} else if strings.ToUpper(valGender) == "F" {
		jnsKelamin = "Perempuan"
	} else {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Invalid Gender"))
		return
	}
	resp.Write([]byte("Method yang digunakan adalah : " + req.Method + ", "))
	resp.Write([]byte("Hello "))
	resp.Write([]byte(name))
	resp.Write([]byte(", Jenis Kelamin Anda Adalah : " + jnsKelamin))

}

func helloHandlerPost(resp http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	// kita harus mengecek error nya
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Println(err.Error())
	}

	resp.Write([]byte("Method yang digunakan adalah : " + req.Method + ", "))
	resp.Write([]byte("isi dari body adalah : "))
	resp.Write(reqBody)

	// cara kedua
	// resp.Write([]byte("isi dari body adalah : " + string(reqBody)))
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	// resp.Write([]byte("Method yang digunakan adalah : " + req.Method + ", "))

	if req.Method == "GET" {
		helloHandlerGet(resp, req)
	} else if req.Method == "POST" {
		helloHandlerPost(resp, req)
	} else {
		// resp.Write dicetak di client
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Not Support Request Method"))
		// dicetak di console
		fmt.Println("Not Support Request Method")
	}
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
