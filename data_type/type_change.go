package data_type

import "fmt"

func TypeChange() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(z)
	fmt.Println(x, y, z, d)
}
