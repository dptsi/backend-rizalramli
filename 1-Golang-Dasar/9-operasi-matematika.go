package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	// Augmented Assignment
	var i = 10
	i+=10
	fmt.Println(i)

	// Unary Operator
	i++
	fmt.Println(i)

	var negatif = -100
	var positive = 100
	fmt.Println(negatif)
	fmt.Println(positive)
}