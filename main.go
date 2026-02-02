package main

import (
	"e-library/handlers"
	"e-library/middleware"
	"e-library/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Use ReleaseMode for cleaner logging and better performance
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Best Practice: Structured Logging + Panic Recovery
	r.Use(middleware.StructuredLogger())
	r.Use(gin.Recovery())

	// Repository Strategy Pattern
	// DEFAULT: MemoryRepo for instant testing
	repo := repository.NewMemoryRepo()

	// UNCOMMENT FOR POSTGRES:
	// db, _ := sql.Open("postgres", "host=localhost user=user password=pass dbname=lib sslmode=disable")
	// repo := repository.NewPostgresRepo(db)

	h := &handlers.LibraryHandler{Repo: repo}

	// Endpoints
	r.GET("/Book", h.GetBook)
	r.POST("/Borrow", h.BorrowBook)
	r.POST("/Extend", h.ExtendLoan)
	r.POST("/Return", h.ReturnBook)

	log.Println("Server listening on :3000")
	err := r.Run(":3000")
	if err != nil {
		log.Println("Server encountered an error.", err)
	}
}
