package main

import "fmt"

func main(){
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	
	// karena batas limit 127 dan akan dimulai dari minus -128 (result minus)
	fmt.Println(nilai8)

	var name = "Ramli"

	// result byte
	var e = name[0]
	// converstion to string
	var eString string = string(e);

	fmt.Println(name)
	fmt.Println(eString)
}