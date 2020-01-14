package main

import (
	"materiGolang/goCleanCode/meja/handler"

	"github.com/gorilla/mux"
)

func main() {
	port := "8000"

	router := mux.NewRouter().StrictSlash(true)

	handler.CreateMejaHandler(router)
}
