package main 

import(
	"fmt"
	go_say_hello "github.com/rizalramli/go-say-hello/v2"
)

func main(){
	// install dependency
	// go mod init 2-Golang-Modules
	// go get github.com/rizalramli/go-say-hello

	// upgrade dependency
	// ubah di file go.mod ganti versi dari v1.0.0 ke v1.1.1
	// lalu ketikan go get

	// jika terjadi major upgrade
	// 1. misal di function ada tambahan parameter
	// 2. hapus require module yang lama
	// 3. go get github.com/rizalramli/go-say-hello/v2

	fmt.Println(go_say_hello.SayHello("Ramli"))
}