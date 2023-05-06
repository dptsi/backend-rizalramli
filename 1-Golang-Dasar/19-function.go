package main

import "fmt"

func sayHello(){
	fmt.Println("Hello")
}

func main(){
	for i:=1; i<10; i++ {
		sayHello()
	}
}