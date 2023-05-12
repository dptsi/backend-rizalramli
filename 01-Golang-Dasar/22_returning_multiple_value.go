package main

import "fmt"

func getFullName() (string,string,int){
	return "Rizal","Ramli", 98
}

func main(){
	// tanda _ artinya di ignore
	firstName,_,born := getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
	fmt.Println(born)
}