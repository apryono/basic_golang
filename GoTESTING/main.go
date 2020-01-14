package main

import (
	"errors"
	"fmt"
)

// HitungLuasPersegi rumus
func HitungLuasPersegi(sisi int) (int, error) {
	if sisi <= 0 {
		return 0, errors.New("error with : %v")
	}
	return sisi + sisi, nil
}

func main() {
	luas, err := HitungLuasPersegi(10)
	if err != nil {
		fmt.Printf("Error with : %v\n", err.Error())
	}
	fmt.Println(luas, err)
}
