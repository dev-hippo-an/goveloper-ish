# Functions
## 기본 특징
- **가변 입력 파라미터를 지원**
    - 대표적으로 fmt.Println() 함수도 가변 파라미터가 적용된 함수이다.
    - 가변 입력 파라미터는 지정된 타입의 슬라이스이다.
    - 호출부에서 슬라이스로 값을 넣어서 호출하는 경우 ... 를 통한 slice unpacking 을 해야 한다.
- **다중 반환 지원**
  - 다중 반환 값을 리턴하는 경우 반환 타입은 소괄호내 쉼표로 구분하여 나열한다.
  - return 구문을 이용한 반환 부분은 괄호 없이 쉼표로 구분한다.
- **return 특성**
  - 함수 호출부에서 반환된 값을 사용하지 않으려면 _ 를 이용하여 명시적으로 무시가 가능하다.
  - 이름이 지정된 return value 혹은 이름이 지정된 반환값의 비어있는 return 은 심각한 런타임 에러를 발생시킬 수 있으니
    - 절대 사용 금지!!

## 함수는 일급 시민
  - 함수는 값이다. -> 이는 함수를 변수로 할당이 가능하다는 뜻이다.
  - 함수의 타입은 키워드(func), 입력 파라미터 타입, 리턴 타입으로 구성되며
  - 이를 **함수 시그니처✔**라고 한다.
  - 동일한 함수 시그니처를 가지는 함수 예시는 다음과 같을 수 있다.
  ```go
  func appendHi(value string) string {
      return fmt.Sprintf("%s hi!", value)
  } 
  
  func appendBye(value string) string {
      return fmt.Sprintf("%s bye!", value)
  }
  ```
  - 함수 내 새로운 함수를 정의해 할당할 수 있다.
  - 이를 **익명 함수** 라고 하며 다음과 같은 형태를 가진다.
  ```go
  for i := 0; i < 5; i++ {
      func(j int) {
          fmt.Println("printing", j, "from inside of an anonymous function")
      }(i)
  }
  
  var anonymousFunction = func() {
      fmt.Println("Anonymous function call")
  }
  
  anonymousFunction()
  ```
## 클로저
- 함수 내부에 선언된 함수가 함수 외부 함수에서 선언한 변수를 접근하고 수정할 수 있는 것을 의미
- 파라미터로 함수를 전달 / 함수에서 함수 반환 등 예시 코드 확인 (functions.closure_test.go)

## defer
- 사용한 자원(보통은 임시 자원으로 파일 입출력, 네트워크, 디비 커넥션 등) 을 정리하도록 하는 키워드
- defer 는 후입 선출의 순서로 실행된다.
  - 마지막에 defer 로 등록된 것이 가장 먼저 실행된다.
- 예시 코드 확인 (functions.defer_test.go)

## call by value
- 고 언어는 **값에 의한 호출**을 사용하는 언어이다.
  - 이는 함수에 파라미터로 넘겨지는 변수는 **항상** 해당 변수의 **복사본** 을 만들어 넘긴다는 의미이다.
  - 