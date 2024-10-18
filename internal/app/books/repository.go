package books

import (
	"context"
	"encoding/json"
	"fiber-template/internal/config"
	"fiber-template/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type BooksRepository interface {
	CreateBook(models.Book) (*models.Book, error)
	GetBookByID(context.Context, int) (*models.Book, error)
	GetBooksByAuthorID(int) ([]models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(int, *models.Book) (*models.Book, error)
	DeleteBook(int) error
}

type RepositoryBooks struct {
	config *config.Config
	db     *gorm.DB
	redis  *redis.Client
	l      zerolog.Logger
}

func NewBooksRepository(cfg *config.Config, db *gorm.DB, redis *redis.Client, log zerolog.Logger) BooksRepository {
	return &RepositoryBooks{
		config: cfg,
		db:     db,
		redis:  redis,
		l:      log,
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

func (r *RepositoryBooks) GetBookByID(ctx context.Context, bookID int) (*models.Book, error) {
	var book models.Book

	cacheKey := "book:" + string(rune(bookID))

	cachedBook, err := r.redis.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		r.l.Info().Int("book_id", bookID).Msg("Fetching book by ID from DB")
		// Si no está en cache, realiza la consulta a la base de datos
		if err := r.db.First(&book, bookID).Error; err != nil {
			return nil, err
		}

		// Guarda el libro en Redis, configurando un tiempo de expiración de cache
		bookJSON, _ := json.Marshal(book)
		err = r.redis.Set(ctx, cacheKey, bookJSON, 10*time.Minute).Err()
		if err != nil {
			return nil, err
		}

		return &book, nil
	} else if err != nil {
		return nil, err
	}

	// Si el libro está en cache, lo deserializamos y lo devolvemos
	if err := json.Unmarshal([]byte(cachedBook), &book); err != nil {
		return nil, err
	}
	r.l.Info().Int("book_id", bookID).Msg("Fetching book by ID from Redis")
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

func (r *RepositoryBooks) UpdateBook(bookID int, book *models.Book) (*models.Book, error) {
	result := r.db.Model(&models.Book{}).Where("book_id = ?", bookID).Updates(book)
	if result.Error != nil {
		return nil, result.Error // Retorna error si ocurrió algún problema
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound // Retorna error si no se encontró el libro
	}

	book.BookID = bookID

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
