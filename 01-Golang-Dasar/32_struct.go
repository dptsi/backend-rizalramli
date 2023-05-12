package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main(){
	var ramli Customer
	ramli.Name = "Ramli"
	ramli.Address = "Lumajang"
	ramli.Age = 24
	fmt.Println(ramli)

	// Short
	ali := Customer {
		Name : "ali",
		Address : "klakah",
		Age : 25,
	}
	fmt.Println(ali)

	kika := Customer{"kika","kunir",30}
	fmt.Println(kika)

}