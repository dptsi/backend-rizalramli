package main 

import "fmt"

type Address struct {
	City,Province,Country string
}

func changeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
	alamat := Address{"Lumajang","Jawa Timur",""}

	var alamatPointer *Address = &alamat
	changeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}