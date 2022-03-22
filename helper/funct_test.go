package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldFaisal(t *testing.T){
	result := HelloWorld("Faisal")

	if result != "Hello Faisal" {
		//unit test fail
		// panic("Result is not 'Hello Faisal'")
		t.Error("Result Must Be 'Hello Faisal'")
	}
	

	fmt.Println("Test Hello World Done")
}

func TestHelloWorldFajar(t *testing.T){
	result := HelloWorld("Fajar")

	if result != "Hello Fajar" {
		//unit test fail
		t.Fatal("Resul Must Be 'Hello Fajar'")
	}
	fmt.Println("Test Hello World Fajar")
}

func TestAddArithmatic(t *testing.T){
	result := AddNumber(1,2)

	if result == 0 {
		panic("Result must be a number")
	}
	fmt.Println("Test Add Arithmmatic Done")
}

func TestHelloWorldAssert(t *testing.T){
	//assert memanggil fail
	result := HelloWorld("Faisal")
	assert.Equal(t, "Hello Faisal", result, "Result Must Be 'Hello Faisal'")
	//assert.Equal(pointer of testing (t), compare, variable value with func, message)
	fmt.Println("TestHelloWorld with assert Done")
}

func TestHelloWorldRequire(t *testing.T){
	//requrie memanggil failnow
	result := HelloWorld("Faisal")
	require.Equal(t, "Hello Faisal", result, "Result Must Be 'Hello Faisal'")
	fmt.Println("TestHelloWorld with reqire Done")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Cannot run on Windows")
	}
	result := HelloWorld("nur")
	require.Equal(t, "Hello Faisal", result, "Result Must Be 'Hello Faisal'")

}

func TestMain(m *testing.M){
	//before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T){
	t.Run("Faisal", func(t *testing.T) {
		result := HelloWorld("Faisal")
		require.Equal(t, "Hello Faisal", result, "Result Must Be 'Hello Faisal'")
	})
	t.Run("Fajar", func(t *testing.T) {
		result := HelloWorld("Nur")
		require.Equal(t, "Hello Fajar", result, "Result Must Be 'Hello Fajar'")
	})
}

func TestHelloWorldTable(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Faisal",
			request: "Faisal",
			expected: "Hello Faisal",
		},
		{
			name: "Fajar",
			request: "Fajar",
			expected: "Hello Fajar",
		},
		{
			name: "Nursaid",
			request: "Nursaid",
			expected: "Hello Nursaid",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0 ; i < b.N; i++{
		HelloWorld("Faisal")
	}
}

//sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Faisal", func(b *testing.B) {
		for i := 0 ; i < b.N; i++ {
			HelloWorld("Faisal")
		}
	})
	b.Run("Fajar", func(b *testing.B) {
		for i := 0 ; i < b.N; i++ {
			HelloWorld("Fajar")
		}
	})
}

//table benchmark

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{name: "Faisal", request: "Faisal"},
		{name: "Fajar", request: "Fajar"},
	}
	for _, benchmark:= range benchmarks{
		b.Run(benchmark.name, func(b *testing.B) {
			for i:= 0; i < b.N ; i++{
				HelloWorld(benchmark.request)
			}
		})
	}

	
}