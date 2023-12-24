package flow_control

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIfStatement(t *testing.T) {
	if n := rand.Intn(111); n == 0 {
		fmt.Println("That's too low")
	} else if n > 100 {
		fmt.Println("that's too big :", n)
	} else {
		fmt.Println("That's a good :", n)
	}

	if x := "asdf"; len(x) < 3 {
		fmt.Println("Length is less than 3 : ", x)
	} else {
		fmt.Println("Length is greater or equal to 3 : ", x)
	}
}
