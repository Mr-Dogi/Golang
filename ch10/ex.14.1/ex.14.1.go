package main

import (
	//"strings"
	"fmt"
)

func ToUpper1(str string) string {
	var rst string
	for _,c := range str {
		if c >= 'a' && c <='z'{
			rst += string('A'+(c-'a'))
		} else {
			rst+= string(c)
		}
	}
	return rst
}

func main(){
	var str string = "Hello world"

	fmt.Println(ToUpper1(str))
}