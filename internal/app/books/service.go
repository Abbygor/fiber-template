package books

import (
	"errors"
	"fiber-template/internal/models"
	"time"
)

type BooksService interface {
	CreateBook(models.Book) (*models.Book, error)
	GetBookByID(int) (*models.Book, error)
	GetBooksByAuthorID(int) ([]models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(int, *models.Book) (*models.Book, error)
	DeleteBook(int) error
}

type ServiceBooks struct {
	bookRepository BooksRepository
}

func NewBooksService(repo BooksRepository) BooksService {
	return &ServiceBooks{
		bookRepository: repo,
	}
}

func (s *ServiceBooks) CreateBook(book models.Book) (*models.Book, error) {
	if book.Title == "" {
		return nil, errors.New("el título del libro es requerido")
	}

	_, err := time.Parse("2006-01-02", book.PublishDate)
	if err != nil {
		return nil, errors.New("formato de fecha de publicación no válido")
	}

	return s.bookRepository.CreateBook(book)
}

func (s *ServiceBooks) GetBookByID(bookID int) (*models.Book, error) {
	book, err := s.bookRepository.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *ServiceBooks) GetBooksByAuthorID(authorID int) ([]models.Book, error) {
	books, err := s.bookRepository.GetBooksByAuthorID(authorID)
	if err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, errors.New("no books found for this author")
	}

	return books, nil
}
func (s *ServiceBooks) GetBooks() ([]models.Book, error) {
	books, err := s.bookRepository.GetBooks()
	if err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, errors.New("no books found")
	}

	return books, nil
}

func (s *ServiceBooks) UpdateBook(bookID int, book *models.Book) (*models.Book, error) {
	updatedBook, err := s.bookRepository.UpdateBook(bookID, book)
	if err != nil {
		return nil, err // Retorna el error si ocurre algún problema
	}
	return updatedBook, nil // Retorna el libro actualizado
}

func (s *ServiceBooks) DeleteBook(bookID int) error {
	err := s.bookRepository.DeleteBook(bookID)
	if err != nil {
		return err // Retorna el error si ocurrió algún problema
	}
	return nil // Retorna nil si la eliminación fue exitosa
}
