package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}

var bookshelf = []Book{
	{
		ID:    1,
		Name:  "Blue Bird",
		Pages: 500,
	},
}
var MAXID int = 1

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "book not found"})
		return
	}

	for _, book := range bookshelf {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the max ID
	for _, book := range bookshelf {
		if book.Name == newBook.Name {
			c.JSON(http.StatusConflict, gin.H{"message": "duplicate book name"})
			return
		}
	}

	// Increment the ID for the new book
	MAXID++
	newBook.ID = MAXID


	// Add the new book to the bookshelf
	bookshelf = append(bookshelf, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, book := range bookshelf {
		if book.ID == id {
			// Delete the book from the bookshelf
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.Status(http.StatusNoContent)
}

func updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook.ID = id
	for _, book := range bookshelf {
		if book.Name == updatedBook.Name && book.ID != id {
			c.JSON(http.StatusConflict, gin.H{"message": "duplicate book name"})
			return
		}
	}

	for i, book := range bookshelf {
		if book.ID == id {
			// Update the book in the bookshelf
			bookshelf[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	// Routes
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	err := r.Run(":8087")
	if err != nil {
		fmt.Println(err)
		return
	}
}
