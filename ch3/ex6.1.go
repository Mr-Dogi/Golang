package main

import "fmt"

func Multiple(a int,b int)int {
	return  a*b
}

func main(){

	a := 3
	b := 5

	fmt.Println(Multiple(a,b))

}