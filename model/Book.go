package model

import "fmt"

type Book struct {
	Name   string
	Writer string
	Pages  int
}

func (book *Book) Create() bool {
	fmt.Println("crete a book")
	return true
}
func (book *Book) Query() {
	fmt.Println("query a book")
}
func (book *Book) Update() {
	fmt.Println("update a book")
}
func (book *Book) Delete() {
	fmt.Println("delete a book")
}
