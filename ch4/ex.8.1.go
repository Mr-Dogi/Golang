package main

import "fmt"

func main(){
	age := 22

	if age<10{
		fmt.Println("You are child")
	}else if age >=20 && age <30 {
		fmt.Println("Best age of life")
	}else {
		fmt.Println("you are beautiful")
	}
}