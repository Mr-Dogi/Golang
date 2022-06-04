package main

import (
	"fmt"
	"strings"
)

func Upper1(s string) string {

	var rst string

	for _, c := range s {
		if c >= 'a' && c <= 'x' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}

	}

	return rst
}

func Upper2(s string) string {

	var builder strings.Builder

	for _, c := range s {
		if c >= 'a' && c <= 'x' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}

	}

	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(Upper1(str))
	fmt.Println(Upper2(str))
}
