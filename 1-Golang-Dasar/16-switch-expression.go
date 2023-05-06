package main

import "fmt"

func main(){
	name := "Ramli"

	switch name {
		case "Ramli":
			fmt.Println("Benar")
		case "Rizal":
			fmt.Println("Salah")
		default :
			fmt.Println("Default")
	}

	// short
	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Benar")
		case false:
			fmt.Println("Salah")
	}

	// tanpa kondisi
	length2 := len(name);
	switch {
		case length2 < 6:
			fmt.Println("Benar")
		case length2 > 6:
			fmt.Println("Salah")
		default:
			fmt.Println("Default")
	}
}