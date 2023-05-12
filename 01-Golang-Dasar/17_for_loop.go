package main

import "fmt"

func main() {
	count := 1;

	for count <= 10 {
		fmt.Println("Perulangan ke ",count)
		count++
	}

	// Pakai Statement
	for count2 := 1; count2 < 10; count2++ {
		fmt.Println("Perulangan ke ",count2)
	}

	// Pakai slice
	slice := []string{"Mohamad","Rizal","Ramli"}
	for i := 0; i < len(slice); i++ {
		fmt.Println("Index ke ",i," = ",slice[i])
	} 

	// Pakai For Range
	for i2, value := range slice {
		fmt.Println("Index ke ",i2," = ",value)
	}

	// Pakai For Range Tanpa Print index
	for _, value := range slice {
		fmt.Println("Value"," = ",value)
	}

	// For Range Dengan Tipe Data Map
	person := make(map[string]string)
	person["name"] = "Ramli"
	person["title"] = "Programmer"
	for key, value := range person {
		fmt.Println(key," = ",value)
	}

}