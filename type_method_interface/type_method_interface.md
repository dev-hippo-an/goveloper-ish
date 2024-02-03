# 타입 / 메서드 / 인터페이스
## 타입
- 구조체를 선언할 때 사용하는 type 은 다음과 같이 선언하고 사용할 수 있습니다.
```go
type Person struct {
    FirstName string
    LastName
}
```
- 위에서 구조체를 선언하는 것은 다음과 같이 해석할 수 있습니다.
  - 구조체 리터럴 (struct) 의 기본 타입을 가지는
  - Person 이란 이름의 사용자가 정의한 타입을 선언
- 고 언어에서는 구조체 리터럴 뿐 만 아니라 기본 타입과 복합 타입 리터럴을 사용하여
- 커스텀한 타입을 정의할 수 있습니다.
```go
type Score int
type Name string  // 기본 타입 리터럴로 커스텀 타입 정의
type Predicate func(string) bool  // 복합 타입 리터럴로 커스텀 타입 정의
type Board map[string]string
```

### 타입 선언의 시점
- 타입은 문서의 개념으로 사용합니다.
  - 개념을 위한 이름을 제공하여 
    - 코드를 명확하게 만들고
    - 기대되는 데이터 종류를 기술합니다.
- 같은 기본 데이터를 가지지만 수행하는 부분이 다르게 사용되는 경우 다른 타입을 만들어 사용할 수 있습니다.
```go
type Money struct{
    Amount float64
}
type Budget Money
type Price Money

```

### 열거형 사용을 위한 iota
- go 언어에서는 열거형 타입을 가지고 있지 않습니다.
- 대신 iota 를 사용해 증가하는 값에 대해 상수 값을 할당할 수 있도록 합니다.
```go
type MovieCategory int

const (
	Uncategorized MovieCategory = iota
	Action
	Fantasy
	Comedy
)
```
- iota 기반의 열거형은 두 가지 한계점이 존재합니다.
  - 사용자가 사용자 타입의 추가적인 값을 생성하는 것을 막을 수 있는 방법이 없습니다.
  - 열거된 순서의 중간에 새로운 식별자가 들어가는 경우 후속된 값이 재계산되기 때문에 기존에 저장되어 있는 값이 있는 경우 디버깅이 어려운 버그가 발생할 수 있습니다.
- **iota 열거를 사용할 때는 상수로 관리가 되기 때문에 사용에 주의가 반드시 필요합니다.**


## 메서드
- 함수 선언과 비슷하지만 리시버를 명시해야 합니다.
- 메서드 선언에 쓰이는 리시버는 func 키워드와 함수명 사이에 작성됩니다.
```go
type Person struct {
    Name string
}

func (p Person) String() string {
    return fmt.Sprintf("My name is %s", p.Name)
}
```


### 포인터 리시버와 값 리시버
```go
func (p *Person) Update(name string) {
    p.Name = name
}
```
- 포인터 파라미터와 동일한 규칙이 적용되며 어떤 타입을 사용하면 좋은지에 대한 가이드는 다음과 같습니다.
  - 메서드가 리시버를 수정한다면 **반드시** 포인터 리시버 사용해야 합니다.
  - nil 인스턴스 처리가 필요한 경우 **반드시** 포인터 리시버 사용해야 합니다.
  - 리시버 수정이 없다면 값 리시버 사용을 고려합니다.
  - 일반적인 관행은 포인터 리시버를 사용하는 메서드가 타입에 있는 경우 
  - 일관성 유지를 위해 모든 메서드에 포인터 리시버를 사용합니다.

```go
person := Person{Name: "king"}
person.Update("kong")
fmt.Println(person.String())
```
- 위의 코드 스니펫의 `person`은 값 타입의 지역 변수이지만 포인터 리시버와 함께 사용할 수 있습니다.
  - go 는 자동으로 지역변수를 포인터 타입으로 변환합니다.
  - `(&person).Update("kong")`

### 함수와 메서드 사용하는 경우
- 함수와 메서드는 매우 유사하며 함수를 사용하는 방식과 비슷하게 메서드를 사용할 수 있습니다.
- 그렇기 때문에 어떤 경우에 함수를 사용하고 메서드를 사용하는지 구분이 필요합니다.
- 구분하는 요소는 **함수가 다른 데이터에 의존적인지 여부**입니다.
  - 시작될 때 설정되거나 프로그램이 수행 중 변경되는 값에 의존하는 경우 구조체에 저장되어야 하며 해당 로직은 메서드로 구현되어야 합니다.
  - 로직이 단지 입력 파라미터에 의존적이라면 함수로 구현할 수 있습니다.
- 타입 - 패키지 - 모듈 - 테스트 - 의존성 주입은 상호 관련된 개념입니다.

### composition 을 위한 임베딩
- go 언어는 상속이란 개념이 없지만 내장된 구성과 승격의 기능을 통해 코드 재사용성을 높일 수 있습니다.
```go
type Employee struct {
    Name string
    ID string
}

func (e Employee) Description() string {
    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
    Employee
    Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// business logic
}

m := Manager {
    Employee: Employee {
        Name: "Sehyeong",
		ID: "12345",
    },
	Reports: []Employee{},
}

fmt.Println(m.ID)  // 12345
fmt.Println(m.Description)  // Sehyeong (12345)
```

- Employee 타입을 임베딩하는 것으로
- 임베딩된 항목의 선언된 모든 항목이나 메서드는 승격(promotion) 되어 구조체를 포함하고 바로 실행도 가능합니다.
- 임베딩은 상속과는 다른 개념입니다.
  - go 언어에서는 코드 실행 중에 객체의 타입에 따라 다른 메서드를 호출할 수 없습니다.
  - 이를 구체 타입을 위한 동적 디스패치를 지원하지 않는다고 할 수 있습니다.
  - 이를 대신해 인터페이스를 사용해서 동적 행동을 구현할 수 있습니다.

## 인터페이스
- go 의 유일한 추상 타입인 암묵적 인터페이스에 대해서 알아보겠습니다.
```go
type Stringer interface {
	String() string
}
```
- 일반적으로 인터페이스는 모든 블록내에 선언이 가능합니다.
- 일반적인 관례에 의하면 인터페이스는 이름의 마지막에 ~er 을 붙입니다.
  - io.Reader, io.ReadCloser, http.Handler 등

### go 에서의 인터페이스
- 다른 언어와 인터페이스의 구현에 있어서 가장 큰 차이점은 다음과 같습니다.
- 인터페이스를 구현하는 구체타입은 구현하는 인터페이스를 선언하지 않고 **암묵적**으로 구현합니다.
- 암묵적 인터페이스 구현의 장점:
  - 코드를 더 간결하고 명확하게 만듭니다. 
  - 구조체가 인터페이스를 구현하는지 여부를 명시적으로 확인할 필요가 없습니다. 
  - 컴파일러가 인터페이스 구현을 자동으로 검사하여 런타임 오류를 방지합니다.
- 암묵적 인터페이스 구현의 단점:
  - 구조체가 인터페이스를 구현하는지 여부를 명시적으로 알 수 없어 코드를 이해하기 어려울 수 있습니다.
  - 인터페이스에 새 메서드를 추가하면 해당 메서드를 구현하지 않는 구조체는 더 이상 인터페이스를 구현하지 않게 됩니다.

- 인터페이스 파라미터를 가지는 함수를 작성하면 코드를 추상화하고 유연하게 만들 수 있습니다.
- 하지만 인터페이스 파라미터에 대해 복사된 메모리는 힙 메모리에 할당되어 성능에 영향을 줄 수 있다.

### 인터페이스의 타입 단언과 타입 스위치
- 인터페이스 변수가 특정 구체 타입을 가지고 있거나 구체 타입이 다른 인터페이스를 구현하는 것을 확인하는 두 가지 방법이 있습니다.
- Type Assertion
  - 인터페이스 변수가 특정 타입인지 확인하고 그 타입의 값을 가져올 수 있게 하는 기능입니다.
  - `x := i.(T)` 형태로 사용됩니다.
  ```go
  
  type Animal interface {
    Speak() string
  }
  
  type Dog struct {}
  
  func (d *Dog) Speak() string {
    return "멍멍!"
  }
  
  type Cat struct {}
  
  func (c *Cat) Speak() string {
    return "야옹!"
  }
  
  func main() {
    var animal Animal = &Dog{}
  
    // Type Assertion을 사용하여 Dog 타입인지 확인하고 값을 가져옴
    dog, ok := animal.(*Dog)
    if ok {
      fmt.Println(dog.Speak()) // "멍멍!" 출력
    } else {
      fmt.Println("Animal은 Dog 타입이 아닙니다.")
    }
  
    // Type Assertion을 사용하여 Cat 타입인지 확인하고 값을 가져옴
    cat, ok := animal.(*Cat)
    if ok {
      fmt.Println(cat.Speak()) // "야옹!" 출력
    } else {
      fmt.Println("Animal은 Cat 타입이 아닙니다.")
    }
  }
  
  ```
  - Type Assertion 사용 시 주의 사항
    - 런타임 오류를 발생시킬 수 있으므로 주의해서 사용해야 합니다. 
    - type assertion 사용 전 인터페이스 변수가 nil 값이 아닌지 확인하는 것이 좋습니다. 
    - Type Switch를 사용하여 인터페이스 변수의 타입을 확인하는 것이 더 안전한 방법일 수 있습니다.
- Type Switch
  ```go
  switch i := i.(type) {
  case T1:
  // T1 케이스 블록
  case T2:
  // T2 케이스 블록
  default:
  // 기본 케이스 블록
  }
  ```
  - 타입 단언보다 더 안정적으로 변수 타입을 확인할 수 있는 방법입니다.
  - 인터페이스에 연관된 타입이 없을 수 있기 때문에 nil 을 사용할 수 있습니다.