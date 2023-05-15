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

func TestMain(m *testing.M){
	// Before
	fmt.Println("before unit test")

	m.Run()

	// After
	fmt.Println("after unit test")
}

func TestSubTest(t *testing.T){
	t.Run("Rizal",func(t *testing.T){
		result := HelloWorld("Rizal")
		require.Equal(t,"Hello Rizal",result,"Harusnya Hello Rizal")
	})

	t.Run("Ramli",func(t *testing.T){
		result := HelloWorld("Ramli")
		require.Equal(t,"Hello Ramli",result,"Harusnya Hello Ramli")
	})
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Mohamad",
			request: "Mohamad",
			expected: "Hello Mohamad",
		},
		{
			name: "Rizal",
			request: "Rizal",
			expected: "Hello Rizal",
		},
		{
			name: "Ramli",
			request: "Ramli",
			expected: "Hello Ramli",
		},
	}

	for _, test := range tests {
		t.Run(test.name,func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t,test.expected,result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i:=0; i<b.N; i++{
		HelloWorld("Ramli")
	}
}

func BenchmarkHelloWorldRamli(b *testing.B) {
	for i:=0; i<b.N; i++{
		HelloWorld("Rizal Ramli")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Rizal",func(b *testing.B){
		for i:=0; i<b.N; i++{
			HelloWorld("Rizal")
		}
	})
	b.Run("Ramli",func(b *testing.B){
		for i:=0; i<b.N; i++{
			HelloWorld("Ramli")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	} {
		{
			name : "Ramli",
			request : "Ramli",
		},
		{
			name : "Rizal",
			request : "Rizal",
		},
	}
	for _,benchmark := range benchmarks {
		for i:=0; i<b.N; i++{
			HelloWorld(benchmark.request)
		}
	}
}