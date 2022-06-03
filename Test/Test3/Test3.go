package main

import "fmt"

func WhatType(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Println(a.(int))
	case string:
		fmt.Println(a.(string))
	case float64:
		fmt.Println(a.(float64))
	default:
		fmt.Println("what the fuck ?")
	}

}

func main() {
	WhatType("1,2,4,5")
	WhatType("Hello")
	WhatType("3.14")
}
