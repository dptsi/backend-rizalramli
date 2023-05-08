package main

import(
	"time"
	"fmt"
)

func main(){
	now:= time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023,05,7,8,30,0,10,time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse,_ := time.Parse(layout,"1998-11-12")
	fmt.Println(parse)
}