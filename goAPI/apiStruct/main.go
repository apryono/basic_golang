package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Student struct {
	ID     int
	Nama   string
	Gender string
}

var students = []Student{
	{101, "Joe", "M"},
	{102, "Maria", "F"},
	{103, "Reza", "M"},
}

func helloHandlerGet(resp http.ResponseWriter, req *http.Request) {

	// FormValue untuk menerima data yang dikirim di url
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

func getAllStudentHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		json, err := json.Marshal(students)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("Opps. Something went wrong. Please contact admin"))
			fmt.Println(err.Error())
			return
		}
		resp.Header().Set("Content-Type", "application/json")
		resp.Write(json)
	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Method tidak di Support"))
		return
	}
}
func studentHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getStudentByID(resp, req)
	} else if req.Method == "POST" {
		insertStudent(resp, req)
	} else if req.Method == "PUT" {
		updateStudent(resp, req)
	} else if req.Method == "DELETE" {
		// deleteStudent(resp, req)
	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Method tidak di Support"))
		return
	}
}

func getStudentByID(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		id := 101
		for _, each := range students {
			if each.ID == id {
				result, err := json.Marshal(each)
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
					resp.Write([]byte("Opps. Something went wrong. Please contact admin"))
					fmt.Println(err.Error())
					return
				}
				resp.Header().Set("Content-Type", "application/json")
				resp.Write(result)
				return
			}
			resp.WriteHeader(http.StatusBadRequest)
			resp.Write([]byte("User Not Found"))
			return

		}
	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Method tidak di Support"))
		return
	}
}

func insertStudent(resp http.ResponseWriter, req *http.Request) {
	// untuk mengambil data dari body
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Println(err.Error())
		return
	}

	var std Student
	err = json.Unmarshal(reqBody, &std)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Opps. Something went wrong. Please contact admin"))
		fmt.Println(err.Error())
		return
	}

	students = append(students, std)
	resp.Write([]byte("Sukses Insert Data !"))
}

func updateStudent(resp http.ResponseWriter, req *http.Request) {
	// reqBody, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	resp.WriteHeader(http.StatusInternalServerError)
	// 	resp.Write([]byte("Opps. Something went wrong. Please contact admin"))
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// id := 101
	// if req.Method == "PUT" {
	// 	for _, each := range students {
	// 		if each.ID == id {
	// 			result, err := json.Marshal(each)
	// 		}
	// 	}
	// }
}

func main() {
	port := "8080"

	http.HandleFunc("/hai", helloHandler)
	http.HandleFunc("/students", getAllStudentHandler)
	http.HandleFunc("/student", studentHandler)
	http.HandleFunc("/student/id", getStudentByID)

	fmt.Println("Starting Web Server At Port : " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
