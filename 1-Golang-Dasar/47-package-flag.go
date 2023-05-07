package main

import(
	"fmt"
	"flag"
)

func main(){
	// name,default value,deskripsi
	var host *string = flag.String("host","localhost","Put your database host")
	var user *string = flag.String("user","root","Put your database user")
	var password *string = flag.String("password","root","Put your database password")

	flag.Parse()
	// return berupa pointer
	// cara run => go run 47-package-flag.go -user=ramli
	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
}