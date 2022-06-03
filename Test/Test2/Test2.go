package main

import "fmt"

func main() {
	var a = []int{1, 2, 4, 5, 6}

	b := a

	b = append(b, 5, 35, 2, 4)

	fmt.Println(b)
	fmt.Println(a)

	c := []int{5, 3, 1, 6, 2}
	var d = make([]int, 5)

	v := copy(d, c)

	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(v)

	var e = []int{4, 6, 1, 6, 1, 3}
	var h = make([]int, 5)

	h = e[:3]

	fmt.Println(h)
	fmt.Println(cap(h))

}
