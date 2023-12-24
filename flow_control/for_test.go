package flow_control

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	fmt.Println("=== 1. 일반 반복문 ===")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=== 2. 조건식만 사용하는 반복문 ===")
	i := 10
	for i < 20 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=== 3. for-range 문 ===")
	evens := []int{2, 4, 6, 8, 10}
	for i, v := range evens {
		fmt.Println(i, v)
	}
}

func TestLabelForLoop(t *testing.T) {
	samples := []string{"hello", "apple!!"}

outer:
	for _, text := range samples {
		for i, r := range text {
			t.Log(i, r, string(r))

			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}
