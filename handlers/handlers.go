package handlers

import (
	"e-library/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LibraryHandler struct {
	Repo repository.LibraryRepository
}

func (h *LibraryHandler) GetBook(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title parameter is required"})
		return
	}
	book, err := h.Repo.GetBook(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *LibraryHandler) BorrowBook(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := h.Repo.BorrowBook(input.Name, input.Title)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, loan)
}

func (h *LibraryHandler) ExtendLoan(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := h.Repo.ExtendLoan(input.Name, input.Title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loan)
}

func (h *LibraryHandler) ReturnBook(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Repo.ReturnBook(input.Name, input.Title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book returned successfully"})
}
