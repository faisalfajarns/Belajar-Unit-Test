package helper

import "testing"

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Faisal")

	if result != "Hello Faisal" {
		//unit test fail
		panic("Result is not 'Hello Faisal'")
	}else {
		
	}
}