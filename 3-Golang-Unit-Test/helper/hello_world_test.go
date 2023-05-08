package helper

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"fmt"
	"runtime"
)

func TestHelloWorldRamli(t *testing.T){
	result := HelloWorld("Ramli")

	if result != "Hello Ramli" {
		// error 
		t.Error("Harusnya Hello Ramli")
	}

	// masih di eksekusi jika error
	fmt.Println("TestHelloWorldRamli di eksekusi")
}

func TestHelloWorldRizal(t *testing.T){
	result := HelloWorld("Rizal")

	if result != "Hello Rizal" {
		// error 
		t.Fatal("Harusnya Hello Rizal")
	}

	// tidak di eksekusi jika error
	fmt.Println("TestHelloWorldRizal di eksekusi")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Ramli")
	assert.Equal(t,"Hello Ramli",result,"Harusnya Hello Ramli")
	fmt.Println("TestHelloWorldAssert di eksekusi")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Ramli")
	require.Equal(t,"Hello Ramli",result,"Harusnya Hello Ramli")
	fmt.Println("TestHelloWorldRequire di eksekusi")
}

func TestSkip(t *testing.T){
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Tidak bisa pakai linux")
	}

	result := HelloWorld("Ramli")
	assert.Equal(t,"Hello Ramli",result,"Harusnya Hello Ramli")
}