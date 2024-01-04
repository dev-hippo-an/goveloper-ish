package functions

import (
	"sort"
	"testing"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func TestClosureParameter(t *testing.T) {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}

	t.Log(people)

	// 파라미터로 넘긴 people 을 함수 안의 함수 (이경우 두번째 파라미터) 에서 접근하고 people 의 순서를 **변경** 할 수 있다.
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	t.Log(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	t.Log(people)
}

func TestClosureFunction(t *testing.T) {
	twoBase := makeMultiple(2)
	threeBase := makeMultiple(3)

	for i := 0; i < 3; i++ {
		t.Log(twoBase(i), threeBase(i))
	}
}

func makeMultiple(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
