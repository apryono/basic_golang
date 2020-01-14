package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MejaHandler struct {
}

var mejaHandler MejaHandler

// pembuatan method
func (m *MejaHandler) getAllMeja(w http.ResponseWriter, r *http.Request) {

}
func (m *MejaHandler) getMejaByID(w http.ResponseWriter, r *http.Request) {

}
func (m *MejaHandler) deleteMeja(w http.ResponseWriter, r *http.Request) {

}
func (m *MejaHandler) insertMeja(w http.ResponseWriter, r *http.Request) {

}

func CreateMejaHandler(r *mux.Router) {

	r.HandleFunc("/tables", mejaHandler.getAllMeja).Methods(http.MethodGet)
	r.HandleFunc("/table", mejaHandler.getMejaByID).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", mejaHandler.deleteMeja).Methods(http.MethodGet)
	r.HandleFunc("/table/{id}", mejaHandler.insertMeja).Methods(http.MethodGet)
}
