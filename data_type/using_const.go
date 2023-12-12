package data_type

import "fmt"

const x int64 = 10
const y = "Hello"

const (
	key  = "id"
	name = "name"
)

func UsingConst() {
	const hello = "hello"
	fmt.Println(x)
	fmt.Println(y)

	//x = x + 1  // Cannot assign to x
	//hello = "hello there"  // Cannot assign to hello
	//
	//fmt.Println(x)
	//fmt.Println(hello)
}
