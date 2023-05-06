package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0;
	for _, value := range numbers {
		total += value
	}
	return total
}

func main(){
	result := sumAll(10,10,10)
	fmt.Println(result)

	// Pakai Slice
	slice := []int {20,30,50}
	result2 := sumAll(slice...)
	fmt.Println(result2)
}