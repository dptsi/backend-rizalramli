package main

import "fmt"

func sayHello(name string) string {
	return "Hello nama saya " +	name;
}

func main(){
	result := sayHello("Ramli")
	fmt.Println(result)
}