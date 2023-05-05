package main

import "fmt"

func main(){
	var nilaiAkhir = 80
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}