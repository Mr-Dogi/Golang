package main

import "fmt"

type Op func(int, int) int

func GetOperator(s string) Op {

	switch s {
	case "+":
		return func(a int, b int) int {
			return a + b
		}
	case "*":
		return func(a int, b int) int {
			return a * b
		}
	default:
		return nil
	}

}

// 익명함수를 반환하는 변수를 만듬

func main() {

	f := GetOperator("+")

	fmt.Println(f(5, 3))

	fn := func(a int, b int) {
		fmt.Println(a + b)
	}
	// 익명함수 선언하고 변수에 저장한다.

	fn(3, 4)

	func(a int, b int) {
		fmt.Println(a / b)
	}(6, 3)
	// 익명함수 선언하여 바로 사용한다.

}
