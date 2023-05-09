package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
	"strconv"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func(){
		time.Sleep(2 * time.Second)
		// mengirim data ke dalam channel
		channel <- "Mohamad Rizal Ramli"
		fmt.Println("Selesai Mengirim data ke channel")
	}()

	// menerima data
	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	// mengirim data ke dalam channel
	channel <- "Mohamad Rizal Ramli"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlySend(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Mohamad Rizal Ramli"
}

func OnlyReceived(channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T){

	channel := make(chan string)
	defer close(channel)

	go OnlySend(channel)
	go OnlyReceived(channel)

	time.Sleep(5 * time.Second)

}

func TestBufferedChannel(t *testing.T){
	channel := make(chan string,3)
	defer close(channel)

	go func(){
		channel <- "Mohamad"
		channel <- "Rizal"
		channel <- "Ramli"
	}()

	go func(){
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func(){
		for i:=0; i< 10; i++{
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data:= range channel {
		fmt.Println("Menerima data",data)
	}

	fmt.Println("Selesai")
}