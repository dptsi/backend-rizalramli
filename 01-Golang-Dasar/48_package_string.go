package main

import(
	"strings"
	"fmt"
)

func main(){
	fmt.Println(strings.Contains("Rizal Ramli","Ramli"))
	fmt.Println(strings.Contains("Rizal Ramli","Eko"))

	fmt.Println(strings.Split("Mohamad Rizal Ramli"," "))

	fmt.Println(strings.ToLower("Mohamad Rizal Ramli"))
	fmt.Println(strings.ToUpper("Mohamad Rizal Ramli"))

	fmt.Println(strings.Trim("aaRizal Ramliaa","a"))
	fmt.Println(strings.ReplaceAll("Mohamad Fadhil Ramli","Fadhil","Rizal"))
}