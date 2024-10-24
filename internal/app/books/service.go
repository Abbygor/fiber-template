package books

import (
	"context"
	"errors"
	"fiber-template/internal/models"
	"time"

	"github.com/rs/zerolog"
)

type BooksService interface {
	CreateBook(models.Book) (*models.Book, error)
	GetBookByID(context.Context, int) (*models.Book, error)
	GetBooksByAuthorID(int) ([]models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(int, *models.Book) (*models.Book, error)
	DeleteBook(int) error
}

type ServiceBooks struct {
	bookRepository BooksRepository
	l              zerolog.Logger
}

func NewBooksService(repo BooksRepository, log zerolog.Logger) BooksService {
	return &ServiceBooks{
		bookRepository: repo,
		l:              log,
	}
}

func (s *ServiceBooks) CreateBook(book models.Book) (*models.Book, error) {
	if book.Title == "" {
		s.l.Error().
			Str("module", "books_service").
			Str("function", "CreateBook").
			Err(errors.New("book title required")).
			Msg("missing field: title")
		return nil, errors.New("book title required")
	}

	_, err := time.Parse("2006-01-02", book.PublishDate)
	if err != nil {
		s.l.Error().
			Str("module", "books_service").
			Str("function", "CreateBook").
			Err(err).
			Msg("publishdate field malformed")
		return nil, errors.New("publishdate field malformed")
	}

	bookCreated, err := s.bookRepository.CreateBook(book)
	if err != nil {
		s.l.Error().
			Str("module", "books_service").
			Str("function", "CreateBook").
			Err(err).
			Msg("Failed to create book in repository")
		return nil, err
	}

	return bookCreated, nil
}

func (s *ServiceBooks) GetBookByID(ctx context.Context, bookID int) (*models.Book, error) {
	book, err := s.bookRepository.GetBookByID(ctx, bookID)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "books_service").
			Str("function", "GetBookByID").
			Int("book_id", bookID).
			Msg("error getting book from repository")
		return nil, err
	}

	return book, nil
}

func (s *ServiceBooks) GetBooksByAuthorID(authorID int) ([]models.Book, error) {
	books, err := s.bookRepository.GetBooksByAuthorID(authorID)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "books_service").
			Str("function", "GetBooksByAuthorID").
			Int("author_id", authorID).
			Msg("error getting book by authorID from repository")
		return nil, err
	}

	if len(books) == 0 {
		s.l.Error().
			Err(errors.New("no books found for this author")).
			Str("module", "books_service").
			Str("function", "GetBooksByAuthorID").
			Int("author_id", authorID).
			Msg("error getting book by authorID from repository")
		return nil, errors.New("no books found for this author")
	}

	return books, nil
}
func (s *ServiceBooks) GetBooks() ([]models.Book, error) {
	books, err := s.bookRepository.GetBooks()
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "books_service").
			Str("function", "GetBooks").
			Msg("error getting all books from repository")
		return nil, err
	}

	if len(books) == 0 {
		s.l.Error().
			Err(errors.New("no books found")).
			Str("module", "books_service").
			Str("function", "GetBooks").
			Msg("no books found")
		return nil, errors.New("no books found")
	}

	return books, nil
}

func (s *ServiceBooks) UpdateBook(bookID int, book *models.Book) (*models.Book, error) {
	updatedBook, err := s.bookRepository.UpdateBook(bookID, book)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "books_service").
			Str("function", "UpdateBook").
			Msg("error updating book from repository")
		return nil, err // Retorna el error si ocurre algún problema
	}
	return updatedBook, nil // Retorna el libro actualizado
}

func (s *ServiceBooks) DeleteBook(bookID int) error {
	err := s.bookRepository.DeleteBook(bookID)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "books_service").
			Str("function", "DeleteBook").
			Msg("error deleting book from repository")
		return err // Retorna el error si ocurrió algún problema
	}
	return nil // Retorna nil si la eliminación fue exitosa
}
