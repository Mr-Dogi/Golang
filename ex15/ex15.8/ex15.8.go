package main

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(st Stringer) {
	r := st.(Reader)
	r.Read()
}

func main() {

}
