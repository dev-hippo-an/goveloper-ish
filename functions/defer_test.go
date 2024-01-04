package functions

import (
	"context"
	"database/sql"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestDefer(t *testing.T) {

	f, err := os.Open(filepath.Base("defer_test.go"))

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			break
		}
	}
}

func TestDeferLIFO(t *testing.T) {
	f1, _ := os.Open(filepath.Base("defer_test.go"))
	f2, _ := os.Open(filepath.Base("defer_test.go"))
	f3, _ := os.Open(filepath.Base("defer_test.go"))

	// third call
	defer func(word string) {
		t.Log("close f1", word)
		f1.Close()
	}("chocolate")

	// second call
	defer func(word string) {
		t.Log("close f2", word)
		f2.Close()
	}("sum")

	// first call
	defer func(word string) {
		t.Log("close f3", word)
		f3.Close()
	}("sing")
}

// 이름이 지정된 반환값을 사용하는 가장 좋은 이유는
// 지연된 함수가 해당 함수의 반환값을 검사하거나 수정하는 경우이다.
// 다음은 defer, 클로저, 이름이 지정된 반환 값을 사용하여 데이터베이스 트랜젝션 정리를 처리한다.
func LazyFuncDefer(ctx context.Context, db *sql.DB, value1, value2 string) (err error) /*이름이 지정된 반환 값 사용*/ {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// defer 와 클로저
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)

	if err != nil {
		return err
	}
	return nil
}
