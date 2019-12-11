package main

import "fmt"

//ini merupakan variable global
var varGlobal = "ini global"

func main() {

	//ini merupakan variabel lokal
	var varLokal = "ini variabel lokal"

	fmt.Println("Helo Scoper")
	fmt.Println(varGlobal)
	fmt.Println(varLokal)
}
