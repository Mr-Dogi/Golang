package  main 

import "fmt"

func main(){
	for i :=2; i < 10; i++{
		if i == 7{
			break
		}
		for j :=1; j <10; j++{
			fmt.Println(i,"*",j,i*j)
		}
		fmt.Println()
	}
}