package repository

import "e-library/models"

type LibraryRepository interface {
	GetBook(title string) (*models.BookDetail, error)
	BorrowBook(name, title string) (*models.LoanDetail, error)
	ExtendLoan(name, title string) (*models.LoanDetail, error)
	ReturnBook(name, title string) error
}
