package main

import "fmt"

func main(){
	const firstName string = "Rizal"
	const lastName = "Ramli"
	const value = 100
	
	// Error 
	// firstName = "Coba Diganti"
	// lastName = "Coba Diganti"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// Multiple declaration
	const (
		firstName2 string = "Rizal"
		lastName2 = "Ramli"
		value2 = 200 
	)
	fmt.Println(firstName2)
	fmt.Println(lastName2)
	fmt.Println(value2)
}