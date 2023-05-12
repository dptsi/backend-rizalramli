package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello(){
	fmt.Println("Hello my name is ",customer.Name)
}

func main(){
	ramli := Customer{"Ramli"}
	ramli.sayHello()
}