package main

import "fmt"

// Defer (di eksekusi meskipun program error)
func logging(){
	fmt.Println("Sukses di eksekusi")
}
func runApplication(value int){
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result",result)
}

// Panic (langsung stop eksekusi code dibawahnya)
func endApp(){
	fmt.Println("Aplikasi selesai")
}
func runApp(error bool){
	defer endApp()
	if(error){
		panic("Aplikasi berhenti")
	}
	fmt.Println("Aplikasi berjalan")
}

// Recover (handle panic agar tetep berjalan dan dimasukan ke error message)
func endApp2(){
	message := recover()
	if message != nil {
		fmt.Println("Aplikasi error dengan pesan :",message)
	}
	fmt.Println("Aplikasi selesai")
}
func runApp2(error bool){
	defer endApp2()
	if(error){
		panic("Aplikasi telah berhenti")
	}
	fmt.Println("Aplikasi berjalan")
}



func main(){
	// defer
	// runApplication(0)

	// panic
	// runApp(true)

	// recover 
	runApp2(true)
}