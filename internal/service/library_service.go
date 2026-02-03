package service

import (
	"e-library-api/internal/models"
	"e-library-api/internal/repository"
	"time"
)

// LibraryServiceInterface defines the behaviors for the library service.
type LibraryServiceInterface interface {
	GetBook(title string) (*models.BookDetail, error)
	BorrowBook(name, title string) (*models.LoanDetail, error)
	ExtendLoan(name, title string) (*models.LoanDetail, error)
	ReturnBook(name, title string) error
}

// LibraryService handles business logic such as 4-week duration for books borrowed and 3-week extension
type LibraryService struct {
	Repo repository.LibraryRepository
}

func NewLibraryService(r repository.LibraryRepository) *LibraryService {
	return &LibraryService{Repo: r}
}

func (s *LibraryService) GetBook(title string) (*models.BookDetail, error) {
	return s.Repo.GetBook(title)
}

func (s *LibraryService) BorrowBook(name, title string) (*models.LoanDetail, error) {
	loan := &models.LoanDetail{
		NameOfBorrower: name,
		BookTitle:      title,
		LoanDate:       time.Now(),
		ReturnDate:     time.Now().AddDate(0, 0, 28), // 4-week rule
	}
	return s.Repo.BorrowBook(loan)
}

func (s *LibraryService) ExtendLoan(name, title string) (*models.LoanDetail, error) {
	loan, err := s.Repo.GetLoan(name, title)
	if err != nil {
		return nil, err
	}

	newReturnDate := loan.ReturnDate.AddDate(0, 0, 21) // 3-week extension rule
	return s.Repo.ExtendLoan(name, title, newReturnDate)
}

func (s *LibraryService) ReturnBook(name, title string) error {
	return s.Repo.ReturnBook(name, title)
}
