package main

import (
	"fmt"
	"materiGolang/goAPI/gomux_middleware/middleware"
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

func helloJoeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Joe"))
	fmt.Println("Hello World")
}

func helloRyoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Ryo"))
	fmt.Println("Hello World")
}

func helloQueryHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	nama := r.FormValue("nama")

	w.Write([]byte("- ID yang dikirim adalah " + id))
	w.Write([]byte("- Nama yang dikirim adalah " + nama))

}

func helloPathHandler(w http.ResponseWriter, req *http.Request) {
	pathVar := mux.Vars(req)

	w.Write([]byte("- ID yang dikirim adalah " + pathVar["id"]))
	w.Write([]byte("- Nama yang dikirim adalah " + pathVar["nama"]))

}

func main() {
	port := "8080"

	// fungsi strictSlash digunakan untuk menghandle "/"
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", helloGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/hello", helloPostHandler).Methods(http.MethodPost)

	// pembuatan Sub Routing
	routerHello := router.PathPrefix("/hello").Subrouter()
	routerHello.HandleFunc("/joe", helloJoeHandler)
	routerHello.HandleFunc("/ryo", helloRyoHandler)
	// END Sub Routing

	router.HandleFunc("/helloQuery", helloQueryHandler).Methods(http.MethodGet)
	router.HandleFunc("/helloPath/{id}/{nama}", helloPathHandler).Methods(http.MethodGet)

	// menempatkan middleware pada semua router nya bersifat global
	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server At Port : " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Opps")
		fmt.Println(err.Error())
	}
}
