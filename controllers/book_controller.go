package controllers

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		log.Printf("Error fetching books: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&book)
	if result.Error != nil {
		log.Printf("Error creating book: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		log.Printf("Error deleting book: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	result := database.DB.First(&book, id)
	if result.Error != nil {
		log.Printf("Error fetching book: %v", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	// Check if book exists
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Validate input
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update book
	result := database.DB.Model(&book).Updates(models.Book{
		Title:         input.Title,
		Author:        input.Author,
		PublishedDate: input.PublishedDate,
		Genre:         input.Genre,
		Description:   input.Description,
		CoverImageUrl: input.CoverImageUrl,
		ISBN:          input.ISBN,
	})

	if result.Error != nil {
		log.Printf("Error updating book: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}
