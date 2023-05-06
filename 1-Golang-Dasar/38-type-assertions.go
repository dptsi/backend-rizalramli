package main

import "fmt"

func random() interface{} {
	return "ini string"
}

func main(){
	var result interface{} = random()
	resultString := result.(string)
	fmt.Println(resultString)  

	// error panic
	// resultInt := result.(int)
	// fmt.Println(resultInt)  

	switch value := result.(type){
		case string: 
			fmt.Println("Value",value,"berupa string")
		case int: 
			fmt.Println("Value",value,"berupa string")
		default:
			fmt.Println("Unknown")
	}
}