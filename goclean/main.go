package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	mejaHandler "goclean/meja/handler"
	"goclean/meja/repo"
	"goclean/meja/usecase"
	menuHandler "goclean/menu/handler"
	"goclean/middleware"
)

func main() {
	port := "8080"
	conStr := "root:P@ssw0rd@tcp(127.0.0.1:3306)/turing"

	db, err := sql.Open("mysql", conStr)
	if err != nil {
		log.Fatal("Error When Connect to DB " + conStr + " : " + err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	mejaRepo := repo.CreateMejaRepoMysqlImpl(db)
	mejaUsecase := usecase.CreateMejaUsecase(mejaRepo)

	mejaHandler.CreateMejaHandler(router, mejaUsecase)
	menuHandler.CreateMenuHandler(router)

	router.Use(middleware.Logger)

	fmt.Println("Starting Web Server at port : " + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}

}
