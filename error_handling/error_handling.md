# 오류 처리
- Go 는 함수 마지막에 error 타입의 값을 반환하는 방식으로 오류를 처리한다.
- Go 는 오류가 반환되는 것을 검출하는 구문인 try catch 같은 구문이 없기 때문에 if 문을 사용하여 오류를 검출해야 한다.
    ```go
    type error interface {
        Error() string
    }
    ```
- 모든 인터페이스 타입에 대한 제로값은 nil 이기 때문에 오류가 발생하지 않았다는 것을 나타내기 위해 오류가 없는 error 반환값에 대해선 nil 을 반환해야 한다.

## Go 에서 에러 생성 방법
- Go 에선 2가지 방법으로 Error 를 생성 가능하다.
  ```go
    err1 := errors.New("this is error message")
    err2 := fmt.Errorf("this is another error message")
  ```
- 오류는 인터페이스이기 때문에 자신만의 커스텀 오류를 만들어 사용할 수 있다.
  ```go
  type Status int
  
  const (
      NotFound Status = iota
      BadRequest
  )
  
  type StatusErr struct {
      Status  Status
      Message string
  }
  
  func (se StatusErr) Error() string {
      return se.Message
  }
  ```

## 오류 래핑 / 언래핑
- 오류 코드에 컨텍스트를 추가하여 확장시키는 것을 오류 래핑이라고 한다.
  ```go
  err := errors.New("error1")
  err2 := fmt.Errorf("wrapping error %w", err)
  ```
- 표준 라이브러리는 오류를 언래핑 하기 위한 함수 errors.Unwrap(err) 함수를 제공한다.
- 이는 래핑된 에러를 언래핑 해주는 아름다운 함수제~~


## 패닉과 복구
- Go 런타임이 다음에 무슨일이 일어날지 알 수 없는 상황에서 패닉을 발생시킨다.
- 패닉이 발생하면 함수는 즉시 종료되며 모든 defer 함수를 실행한다.
- Go 는 recover 내장 함수로 패닉을 확인해 안정적인 종료를 제공하고 종료를 방지수 있도록 해준다.
- recover 는 반드시 defer 함수 내에서 호출해야 한다.
  - panic 이 발생하면 defer 로 등록된 함수만 실행할 수 있기 때문이다.