package go_common_pattern

import (
	"testing"
)

type Book struct {
	Name  string
	Price int
	Genre Genre
}

func NewBook(name string, price int, genre Genre) Book {
	return Book{
		Name:  name,
		Price: price,
		Genre: genre,
	}
}

func TestConstructor(t *testing.T) {
	book := NewBook("Harry Potter", 10, Fantasy)
	t.Logf("%+v\n", book)
}

type Genre int64

const (
	Action Genre = iota
	Romance
	Horror
	Fantasy
	Tech
)

// 실제 enum 값은 숫자 값으로 들어간다.
func TestEnum(t *testing.T) {
	book1 := NewBook("Harry Potter", 10, Fantasy)
	book2 := NewBook("Effective Java", 22, Tech)
	t.Logf("%+v\n", book1)
	t.Logf("%+v\n", book2)
}
