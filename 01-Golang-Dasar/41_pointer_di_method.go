package main 

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	ramli := Man{"Ramli"}
	ramli.Married()
	fmt.Println(ramli.Name)
}