package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func main() {

	// untuk memanggil struct nya person untuk referensi di gunakan di variabel di person1
	var person1 person
	person1.name = "Ryo"
	person1.age = 23
	person1.gender = "Male"

	fmt.Println(person1.name)
}
