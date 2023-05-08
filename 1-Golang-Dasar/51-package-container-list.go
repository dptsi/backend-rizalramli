package main

import(
	"container/list"
	"fmt"
)

func main(){
	data := list.New()
	data.PushBack("Mohamad")
	data.PushBack("Rizal")
	data.PushBack("Ramli")

	data.PushFront("Bima")

	// dari depan ke belakang
	// for element := data.Front(); element != nil; element = element.Next(){
	// 	fmt.Println(element.Value)
	// }

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev(){
		fmt.Println(element.Value)
	}
}