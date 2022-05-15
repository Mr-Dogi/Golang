package main

import "fmt"

type Actor struct{
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, HP int, speed float64) *Actor {
	actor := Actor{name,HP,speed}

	return &actor

}

func main(){
	var actor = NewActor("금도끼",99,100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}