package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
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
	{ Id: 1, AuthorId: 1, Title: "War and Peace", Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam" },
	{ Id: 2, AuthorId: 1, Title: "Childhood", Description: "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit" },
	{ Id: 3, AuthorId: 2, Title: "Crime and Punishment", Description: "The plot centers around Raskolnikov’s plan to murder a pawnbroker for her money. He believes this act is justifiable to escape poverty and use the stolen money for good deeds. But it’s not the crime itself that’s the focus of this book, but the aftermath of it." },
  { Id: 4, AuthorId: 2, Title: "The Brothers Karamazov", Description: 
		"Imagine yourself in a small Russian town, getting to know the Karamazov family.There’s the father, Fyodor Pavlovich, a morally dubious figure, and his three sons—the intellectual Ivan, the passionate Dmitri, and the saintly Alyosha." },
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HealthInfo{Healthy: "true"})
}

func getBooks(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, books)
}

func getBookInfo(c *gin.Context) {
	id_str := c.Param("id")

	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uncorrect id"})
		return
	}

	book_id, err := strconv.Atoi(id_str)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uncorrect id, not a number"})
		return
	}

	for _, v := range books {
		if v.Id == book_id {
			c.JSON(http.StatusOK, v)	
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Book with requested id not found"})
}

func getAuthorInfo(c *gin.Context) {
	id_str := c.Param("id")

	if id_str == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uncorrect id"})
		return
	}

	author_id, err := strconv.Atoi(id_str)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uncorrect id, not a number"})
		return
	}

	for _, v := range authors {
		if v.Id == author_id {
			c.JSON(http.StatusOK, v)	
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Book with requested id not found"})
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

	router.GET("/books/:id", getBookInfo)
	router.GET("/authors/:id", getAuthorInfo)

  router.Run(address)
}
