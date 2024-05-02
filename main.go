package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//To store our book we are going to use a struct

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity	"`
}

var books = []book{
	{ID: "1", Title: "In Search Of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookId(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found!")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBooks) //we'll be using the same endpoint
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}
