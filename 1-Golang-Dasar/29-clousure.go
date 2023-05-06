package main 

import "fmt"

func main(){
	counter := 0
	name := "ramli"

	increment := func() {
		name := "rizal"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	} 

	increment()

	fmt.Println(name)
}