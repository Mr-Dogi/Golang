package main

import "fmt"

type Direction int

const(
	None Direction = iota
	North
	East
	South
	West 
)

func DirectionToString(d Direction) string {
	switch d{
	case North:
		fmt.Println(d)
		return "North"
	case East:
		fmt.Println(d)
		return "East"
	case West:
		fmt.Println(d)
		return "West"
	case South:
		fmt.Println(d)
		return "South"
	default :
		fmt.Println(d)
		return "none"
	}
}

func GetDirection(angel float64) Direction {
	switch {									//조건식에는 비굣값을 적지 않는다.
	case angel>=315 , angel >=0 && angel <45:
		return North
	case angel >= 45 && angel<135:
		return East
	case angel>=135 && angel<225:
		return South
	case angel>=255 && angel<=315:
		return West
	default :
		return None
	}
}

//조건을 비교 후 상수를 반환함

func main(){

	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.3)))
	fmt.Println(DirectionToString(GetDirection(94.3)))
	fmt.Println(DirectionToString(GetDirection(-30)))

}

