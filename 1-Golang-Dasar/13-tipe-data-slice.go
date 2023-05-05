package main

import "fmt"

func main(){
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7];
	fmt.Println(slice1)

	// Function slice
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2,"Ramli")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	
	fmt.Println(slice2)
	fmt.Println(months)

	// Buat slice tanpa extends ke arary
	newSlice := make([]string,2,5)
	newSlice[0] = "Rizal"
	newSlice[1] = "Ramli"
	fmt.Println(newSlice)

	// Copy Slice
	copySlice := make([]string,len(newSlice),cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	// array vs slice
	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}