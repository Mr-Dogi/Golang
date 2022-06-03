package main

import "fmt"

type Fc func(int, int) int

func add(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a / b
}

func GetOperator(s string) Fc {
	switch s {
	case "-":
		return add
	case "/":
		return mul
	default:
		return nil
	}
}

func main() {

	var TestMen Fc

	TestMen = GetOperator("/")

	var hi = TestMen(6, 3)

	fmt.Println(hi)

}
