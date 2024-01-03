package functions

import (
	"fmt"
	"testing"
)

func printIntValues(values ...int) {
	fmt.Println(values)
}

func TestVariadicParameters(t *testing.T) {
	printIntValues(1, 2, 3, 4, 5)
	printIntValues([]int{6, 7, 8, 9, 0}...)
}

func TestSignature(t *testing.T) {
	var sig func(string) string

	sig = appendHi
	t.Log(sig("sehyeong"))

	sig = appendBye
	t.Log(sig("fat"))

	// 타입 키워드를 사용해 함수 타입을 정의할 수 있다.
	var app Append
	app = appendHi
	t.Log(app("ejejc"))

}

func TestAnonymous(t *testing.T) {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}

	var anonymousFunction = func() {
		fmt.Println("Anonymous function call")
	}

	anonymousFunction()
}

type Append func(string) string

func appendHi(value string) string {
	return fmt.Sprintf("%s hi!", value)
}

func appendBye(value string) string {
	return fmt.Sprintf("%s bye!", value)
}
