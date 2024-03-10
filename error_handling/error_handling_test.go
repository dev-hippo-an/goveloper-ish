package error_handling

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorHandling(t *testing.T) {
	err := errors.New("this is error message")

	fmt.Println(err) // Error() 메서드가 호출됨

	err = fmt.Errorf("this is another error message")

	fmt.Println(err) // Error() 메서드가 호출됨
}

func TestErrorWrapping(t *testing.T) {
	err := fmt.Errorf("error while testing")

	wrapError := fmt.Errorf("wrap error %w", err)

	fmt.Println(wrapError)
	fmt.Println(errors.Unwrap(wrapError))
}

// 상수로 상태코드를 나타내어 오류를 정의하고 비교할 수 있다.
type Status int

const (
	NotFound Status = iota
	BadRequest
)

type StatusErr struct {
	Status  Status
	Message string
	err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.err
}
