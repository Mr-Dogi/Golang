package main

import "fmt"

var g int = 10      // 전역변수선언

func main(){

	var m int = 20

	{
		var s int = 50
		fmt.Println(m ,s , g)
	}

	// m = s+20

	// 지역변수의 생존은 {} 안에서만 생존하고 끝나면 할당이 해제된다

}