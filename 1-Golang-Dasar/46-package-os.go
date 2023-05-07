package main

import(
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args)

	hostname,err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error dengan pesan",err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}