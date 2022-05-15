package main

import "fmt"

func main(){

	i := 0

	for  i < 10  {
		i++
		if i ==3 {
			continue
		}

		if i == 6{
			break
		}

		fmt.Println("6*",i,"=",6*i)
	}
}