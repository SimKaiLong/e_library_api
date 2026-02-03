package main

import (
	"e-library-api/internal/handlers"
	"e-library-api/internal/middleware"
	"e-library-api/internal/repository"
	"e-library-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Use ReleaseMode for cleaner logging and better performance
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Best Practice: Structured Logging + Panic Recovery
	r.Use(middleware.StructuredLogger())
	r.Use(gin.Recovery())

	// Service Pattern
	var repo repository.LibraryRepository
	var svc service.LibraryServiceInterface

	// DEFAULT: MemoryRepo for instant testing
	repo = repository.NewMemoryRepo()
	svc = service.NewLibraryService(repo)

	// UNCOMMENT FOR POSTGRES:
	//db, err := sql.Open("postgres", "host=localhost user=user password=pass dbname=lib sslmode=disable")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//// Best Practice: Connection Pool Settings
	//db.SetMaxOpenConns(25)
	//db.SetMaxIdleConns(25)
	//db.SetConnMaxLifetime(5 * time.Minute)
	//
	//repo = repository.NewPostgresRepo(db)
	//svc = service.NewLibraryService(repo)

	h := &handlers.LibraryHandler{Service: svc}

	// Endpoints
	r.GET("/Book", h.GetBook)
	r.POST("/Borrow", h.BorrowBook)
	r.POST("/Extend", h.ExtendLoan)
	r.POST("/Return", h.ReturnBook)

	log.Println("Server listening on :3000")
	if err := r.Run(":3000"); err != nil {
		log.Println("Server encountered an error.", err)
	}
}
