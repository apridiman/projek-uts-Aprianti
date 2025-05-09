package handlers

import (
	"library-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BorrowBook(c *gin.Context) {
	var borrowing models.Borrowing
	if err := c.ShouldBindJSON(&borrowing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.First(&book, borrowing.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Decrease available book quantity
	if book.QuantityAvailable <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No available books"})
		return
	}

	book.QuantityAvailable -= 1
	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	borrowing.BorrowDate = time.Now()
	borrowing.Status = "borrowed"
	if err := db.Create(&borrowing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowing)
}

func ReturnBook(c *gin.Context) {
	var borrowing models.Borrowing
	if err := c.ShouldBindJSON(&borrowing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.First(&borrowing, borrowing.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrowing record not found"})
		return
	}

	// Update return date and status
	now := time.Now()
	borrowing.ReturnDate = &now
	borrowing.Status = "returned"
	if err := db.Save(&borrowing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Increase available book quantity
	var book models.Book
	if err := db.First(&book, borrowing.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	book.QuantityAvailable += 1
	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowing)
}

func ListBorrowings(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var borrowings []models.Borrowing
	if err := db.Preload("Book").Preload("Member").Find(&borrowings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, borrowings)
}
