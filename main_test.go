package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"e-library/handlers"
	"e-library/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *repository.MemoryRepo) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	repo := repository.NewMemoryRepo()
	h := &handlers.LibraryHandler{Repo: repo}

	r.GET("/Book", h.GetBook)
	r.POST("/Borrow", h.BorrowBook)
	r.POST("/Extend", h.ExtendLoan)
	r.POST("/Return", h.ReturnBook)

	return r, repo
}

func TestEndpoints(t *testing.T) {
	router, repo := setupTestRouter()

	t.Run("GET /Book - Found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/Book?title=Clean+Code", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /Book - Not Found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/Book?title=Unknown", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("POST /Borrow - Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		payload, _ := json.Marshal(map[string]string{"name": "Alice", "title": "Clean Code"})
		req, _ := http.NewRequest("POST", "/Borrow", bytes.NewBuffer(payload))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("POST /Borrow - Conflict (Out of Stock)", func(t *testing.T) {
		repo.Books["Design Patterns"].AvailableCopies = 0
		w := httptest.NewRecorder()
		payload, _ := json.Marshal(map[string]string{"name": "Bob", "title": "Design Patterns"})
		req, _ := http.NewRequest("POST", "/Borrow", bytes.NewBuffer(payload))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusConflict, w.Code)
	})
}
