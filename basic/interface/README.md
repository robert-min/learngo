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

## 추상화
* 내부 동작을 감춰서 서비스 제공자와 사용자 모두에게 자유를 주는 방식

## Duck Typing
* 만약 새를 봤는데 그 새가 오리처럼 걷고 오리처럼 날고 오리처럼 소리를 내면 그 새를 오리라고 부르겠다. => Go, Python은 Duck Typing을 제공하는 언어
* 타입 선언 시 인터페이스 구현 여부를 명시하지 않아도 됨
* 사용자 중심의 코딩이 가능(서비스 제공자는 구체화된 객체를 제공하고 사용자가 필요에 따라 인터페이스를 정의해서 사용할 수 있음)




[checkpoint](https://youtu.be/IV6zYG3GY5s)

