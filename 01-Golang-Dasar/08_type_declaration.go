package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool

	var ktpRamli NoKTP = "111111"
	var marriedStatus Married = true
	fmt.Println(ktpRamli)
	fmt.Println(marriedStatus)
}