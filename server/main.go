package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

const address string = "localhost:8080"

type HealthInfo struct {
	Healthy   string   `json:"healthy"`
}

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

var nextAuthorId int
var nextBookId int

var authors = []Author {
	{ Id: 1, FirstName: "Lev", LastName: "Tolstoy" },
	{ Id: 2, FirstName: "Fyodor", LastName: "Dostoevsky" },
}

var books = []Book {
	{ Id: 1, AuthorId: 1, Title: "War and Peace", Description: "One of famous russian book, very big" },
	{ Id: 2, AuthorId: 1, Title: "Childhood", Description: "Very interesting book" },
	{ Id: 3, AuthorId: 2, Title: "Crime and Punishment", Description: "Book about student in St.Petersburg" },
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HealthInfo{Healthy: "true"})
}

func getBooks(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, books)
}

func getAuthors(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, authors)
}

func CORSinit(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	c.Next()
}

func main() {
	nextAuthorId = len(authors)
	nextBookId = len(books)

  router := gin.Default()

	router.Use(CORSinit)

	router.GET("/health", getHealth)

	router.GET("/books", getBooks)
	router.GET("/authors", getAuthors)

  router.Run(address)
}
