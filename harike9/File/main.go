package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type student struct {
	name string
	age  int
}

var students = []student{
	{name: "ryo", age: 23},
	{name: "agus", age: 24},
}

var path = "C:/Users/Lenovo/go/src/materiGolang/harike9/File/resource/form-mahasiswa.txt"

func main() {
	writeToFile()
	// readFile()
	// readScanner()

}

func readFile() {
	var text = make([]byte, 100)
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err == nil {
		for {
			isi, _ := file.Read(text)
			if isi == 0 {
				break
			}
			fmt.Println(string(text))
		}
	}

}

func writeToFile() {
	// var file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644) // untuk membaca dan menulis ulang atau create
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644) // untuk menambah file atau create file
	if err != nil {
		log.Fatal(err)
	}
	for _, student := range students {
		/*
			fix-length untuk mensejajarkan nama
			untuk %v nya bisa di tambah atau diatur dengan %10v : artinya posisiny mulai dari kanan
			untuk %v nya bisa di tambah atau diatur dengan %-10v : artinya posisiny mulai dari kanan
		*/
		strStudent := fmt.Sprintf("%-7v, %v\n", student.name, student.age)
		file.WriteString(strStudent)
	}
	file.Sync()

}

func readScanner() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			runes := scanner.Text()
			fmt.Println(runes)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
