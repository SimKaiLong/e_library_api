package repository

import (
	"e-library-api/internal/models"
	"time"
)

type LibraryRepository interface {
	GetBook(title string) (*models.BookDetail, error)
	GetLoan(name, title string) (*models.LoanDetail, error)
	BorrowBook(loan *models.LoanDetail) (*models.LoanDetail, error)
	ExtendLoan(name, title string, newReturnDate time.Time) (*models.LoanDetail, error)
	ReturnBook(name, title string) error
	Ping() error
}
