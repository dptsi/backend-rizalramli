package main 

import(
	"fmt"
	"strconv"
)

func main(){
	boolean,error := strconv.ParseBool("true")
	if error == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(error.Error())
	}

	// angka, desimal (oktal,hexa), int64
	number,error := strconv.ParseInt("1000000",10,64)
	if error == nil {
		fmt.Println(number)
	} else {
		fmt.Println(error.Error())
	}

	value := strconv.FormatInt(100000,10)
	fmt.Println(value)

	// tanpa handle error nill
	valueInt,_ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}