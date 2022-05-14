package main

import "fmt"

func main(){
	a:=3
	var b float64 = 3.5
    
	var c int = int(b)
	d := float64(a *c)               // int를 float로 변환

	var e int64 = 7
    f := int64(d) * e                // float을 int로 변환

	var g int = int(b * 3)           // float을 int로 변환
	var h int = int(b) * 3           // float을 int로 변환

	fmt.Println(g, h , f)

}