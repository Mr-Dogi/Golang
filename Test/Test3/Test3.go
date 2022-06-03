package main

import "fmt"

func WhatType(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Println(a.(int))
		fmt.Printf("%T \n", a)
	case string:
		fmt.Println(a.(string))
		fmt.Printf("%T \n", a)
	case float64:
		fmt.Println(a.(float64))
		fmt.Printf("%T \n", a)
	default:
		fmt.Println("what the fuck ?")
	}

}

func main() {
	WhatType(1)
	WhatType("Hello")
	WhatType(3.14)
}
