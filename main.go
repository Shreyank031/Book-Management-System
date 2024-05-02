package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct called "book" to store information about each book
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity	"`
}

// getBooks returns the list of all books
var books = []book{
	{ID: "1", Title: "In Search Of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// getBooks returns the list of all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// createBooks adds a new book to the list
func createBooks(c *gin.Context) {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// bookById retrieves a book by its ID
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// getBookById returns a book by its ID
func getBookId(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found!")
}

// checkOutBook allows a user to check out a book
func checkOutBook(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBookId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// returnBook allows a user to return a book
func returnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBookId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)     //get req-> getting information
	router.POST("/books", createBooks) //post -> adding/creating new information
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkOutBook) //patch -> updating
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080") // Start the server on localhost:8080
}
