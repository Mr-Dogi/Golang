package main

import _ "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

//func (f *File) Close() {
//
//}
// FILE이 Close 메서드를 포함하고 있지않아
// reader인터페이스를 close 인터페이스로 변환하여도 close가 없어 사용하지 못한다.
// 하지만 file에 close가 선언되어있다면 사용가능하다다
func ReadFile(reader Reader) {
	c := reader.(Closer)
	c.Close()

}

func main() {
	file := &File{}

	ReadFile(file)
}
