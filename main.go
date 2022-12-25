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
