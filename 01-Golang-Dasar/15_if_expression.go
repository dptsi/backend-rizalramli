package main

import "fmt"

func main(){
	name := "Rizal"

	if name == "Ramli" {
		fmt.Println(true)
	} else if name == "Rizal" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Sudah benar")
	}
}