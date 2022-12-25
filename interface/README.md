# 인터페이스
* 구체화된 객체가 아닌 추상화된 상호작용으로 관계를 표현
* 구현이 포함되지 않은 관계만을 나타낸 것
* 인터페이스도 타입


```go
type Interface interface {
    // 구현이 빠진 method
    Fly()
    Walk(distance int) int
}

```

## 인터페이스 규칙
1. 매서드는 만드시 메서드 명이 있어야 함
2. 매개변수와 반환이 다르더라도 이름이 같은 매서드는 있을 수 없음
3. 인터페이스에서는 메서드 구현을 포함하지 않아야 함.

```go
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"민준", 14}

	var stringer Stringer = student

	fmt.Printf("%s\n", stringer.String())
}

```

[checkpoint](https://youtu.be/57Ea64-Nf2U?t=2101)

