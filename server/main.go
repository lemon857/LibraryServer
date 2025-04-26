package main

import (
	"fmt"
)

type Author struct {
	Id        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
}

type Book struct {
	Id          int    `json:"id"`
	AuthorId    int    `json:"authorId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Hello, World!")
}
