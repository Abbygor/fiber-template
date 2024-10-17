package books

import (
	"fiber-template/internal/config"
	"fiber-template/internal/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BooksRepository interface {
	CreateBook(models.Book) (*models.Book, error)
	GetBookByID(int) (*models.Book, error)
	GetBooksByAuthorID(int) ([]models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(*models.Book) (*models.Book, error)
	DeleteBook(int) error
}

type RepositoryBooks struct {
	config *config.Config
	db     *gorm.DB
	redis  *redis.Client
}

func NewBooksRepository(cfg *config.Config, db *gorm.DB, redis *redis.Client) BooksRepository {
	return &RepositoryBooks{
		config: cfg,
		db:     db,
		redis:  redis,
	}
}

func (r *RepositoryBooks) CreateBook(book models.Book) (*models.Book, error) {
	// Crear el libro en la base de datos
	if err := r.db.Create(&book).Error; err != nil {
		return nil, err
	}

	// Devolver el libro creado
	return &book, nil
}

func (r *RepositoryBooks) GetBookByID(bookID int) (*models.Book, error) {
	var book models.Book

	if err := r.db.First(&book, bookID).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
func (r *RepositoryBooks) GetBooksByAuthorID(authorID int) ([]models.Book, error) {
	var books []models.Book

	if err := r.db.Where("author_id = ?", authorID).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil

}

func (r *RepositoryBooks) GetBooks() ([]models.Book, error) {
	var books []models.Book

	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *RepositoryBooks) UpdateBook(book *models.Book) (*models.Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *RepositoryBooks) DeleteBook(bookID int) error {
	result := r.db.Delete(&models.Book{}, bookID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
