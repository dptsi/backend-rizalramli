package main

import (
	"1-Golang-Dasar/helper"
	"fmt"
)

func main(){
	// Tidak bisa diakses ke luar karena nama function huruf kecil
	// helper.sayGoodBye("Ramli")
	// fmt.Println(helper.application)

	fmt.Println(helper.Version)
}