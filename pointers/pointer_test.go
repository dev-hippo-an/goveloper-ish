package pointers

import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {
	var pointerX *string // zero is nil

	fmt.Println(pointerX)
	val := "good morning"
	pointerX = &val

	fmt.Println("this is address:", pointerX)
	fmt.Println("this is value:", *pointerX)
}

func TestNilPointerReassign(t *testing.T) {
	var f *int
	fmt.Println(f)
	failedUpdate(f)
	fmt.Println(f)

	x := 555
	f = &x

	failedUpdate(f)
	fmt.Println(*f)
	update(f)
	fmt.Println(*f)
}

func update(px *int) {
	*px = 200
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}
