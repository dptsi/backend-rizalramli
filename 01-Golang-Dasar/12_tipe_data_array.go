package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "Mohamad"
	names[1] = "Rizal"
	names[2] = "Ramli"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// short
	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

	// function
	fmt.Println(len(names))
	fmt.Println(len(values))
}