package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[897] = Product{"샤프심", 500}
	m[345] = Product{"샤프", 3000}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println()

	var SortStr = make([]int, 0, len(m))

	for k := range m {
		fmt.Println("k", k)
		SortStr = append(SortStr, k)
	}

	sort.Ints(SortStr)

	for _, k := range SortStr {
		fmt.Println(k, m[k])
	}
}
