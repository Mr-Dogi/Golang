package main

import "fmt"

type Stringer interface {
	String() string // String() 메서드를 가지고 있는
	// 구조체는 자동으로 덕 타이핑된다.
}

type Student struct {
	Age int
}

func (s Student) String() string { // 해당 메서드가 String() 가지고 있다.
	// 즉 Stringer 인터페이스에 포함된다. 자동으로
	return fmt.Sprintf("안녕! 나는 %d살이야", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(Student) // 인터페이스 타입에서
	fmt.Println(s.Age)
}

func main() {

	s := Student{13}

	PrintAge(s)

}
