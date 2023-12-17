package composite_type

import (
	"fmt"
	"testing"
)

func TestArrayAndSlice(t *testing.T) {
	var array [10]int
	var slice []int

	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(slice == nil) // should be true

	var compareArray [10]int
	//var compareSlice []int

	fmt.Println(array == compareArray) // should be true
	//fmt.Println(slice == compareSlice) // should be compile error

	fmt.Println(len(slice)) // nil should have size 0
	fmt.Println(len(array)) // size of 10
	var x []int
	x = append(x, 10)
	x = append(x, 1, 2, 3, 5)

	var y = []int{6, 7, 8}
	x = append(x, y...)
	fmt.Println(x) // {10, 1, 2, 3, 5, 6, 7, 8}

	pet := struct {
		name string
		kind string
	}{
		name: "Merry",
		kind: "dog",
	}

	fmt.Println(pet)
}
