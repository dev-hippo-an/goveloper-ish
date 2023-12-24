package flow_control

import (
	"fmt"
	"testing"
)

func TestShadowingVariable(t *testing.T) {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func TestMultipleShadowingVariable(t *testing.T) {
	x := 10

	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}

	fmt.Println(x)
}

func TestPackageNameShadowing(t *testing.T) {
	x := 10
	fmt.Println(x)
	//fmt := "oh no!!!!!!!!!!!!" // 이후 fmt 패키지를 사용할 수 없게 된다.
	//fmt.Println(fmt)
}
