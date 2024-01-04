package functions

import (
	"testing"
)

// struct 도 일반 변수와 같이 값으로 취급된다.
func TestCallByValueStruct(t *testing.T) {
	person := Person{
		FirstName: "sehyeong",
		LastName:  "An",
		Age:       30,
	}

	modPerson(person)
	t.Log(person)
}

func modPerson(p Person) {
	p.Age = 40
	p.LastName = "Power"
}

// map 과 slice 는 포인터로 구현되어 있기 때문에 값의 변경이 일어난다.
func TestCallByValueMapSlice(t *testing.T) {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	t.Log(m)

	s := []int{1, 2, 3}
	modSlice(s)
	t.Log(s)
}

// map 의 모든 변경은 기존 변수에 반영된다.
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

// slice 의 모든 요소의 변경은 가능하나 길이를 늘이는 것은 안된다.
func modSlice(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
	s = append(s, 10)
}
