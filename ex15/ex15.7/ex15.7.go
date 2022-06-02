package main

import _ "fmt"

type AwesomeDB interface {
	GetDate()
	WriteData()
	Close()
}

type OurDB struct {
}

func (db *OurDB) GetData() {

}

func (db *OurDB) WriteData() {

}

func (db *OurDB) Close() {

}

func main() {

}
