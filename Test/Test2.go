package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type Close interface {
	Close()
}

type ReadOnly struct {
}

func (R ReadOnly) Close() {
	fmt.Println("Read")
}

type WriteOnly struct {
}

func (R WriteOnly) Close() {
	fmt.Println("wirte")
}

func Work(cl Close) {
	TestMen := cl.(ReadOnly)

	TestMen.Close()
}

func main() {

	ReadOnly := ReadOnly{}
	Work(ReadOnly)

}
