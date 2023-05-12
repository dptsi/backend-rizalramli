package main

import "fmt"

func getFullName() (firstName string,lastName string){
	firstName = "Rizal"
	lastName = "Ramli"

	return
	// bisa pakai `return firstName,LastName``
}

func main(){
	a,b := getFullName();
	fmt.Println(a)
	fmt.Println(b)
}