package main

import "fmt"

type school struct {
	name       string
	schoolType string
	address    string
}

type person struct {
	name       string
	age        int
	gender     string
	schoolType school
}

func main() {

	var data person
	// cara 1 untuk memanggil struct
	/*
		data.name = "andi"
		data.age = 23
		data.gender = "male"
		data.schoolType.name = "Budi Mulia"
		data.schoolType.schoolType = "BootCamp"
		data.schoolType.address = "Medan"
	*/

	//cara 2 memanggil struct

	fmt.Println(data)

}
