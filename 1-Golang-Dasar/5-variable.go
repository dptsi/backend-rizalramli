package main

import "fmt"

func main(){
	var name string
	
	name = "Rizal Ramli"
	fmt.Println(name)

	var name2 = "Rizkika Zakka"
	fmt.Println(name2)

	var age uint = 23
	fmt.Println(age)

	// short variable
	country := "Indonesia"
	fmt.Println(country)

	// Multiple variable
	var (
		firstName = "Rizal"
		lastName = "Ramli"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}