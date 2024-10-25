package authors

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

type AuthorsRepository interface {
	CreateAuthor(models.Author) (*models.Author, error)
	GetAuthorByID(context.Context, int) (*models.Author, error)
	GetAuthors() ([]models.Author, error)
	UpdateAuthor(int, *models.Author) (*models.Author, error)
	DeleteAuthor(int) error
}

type RepositoryAuthors struct {
	config *config.Config
	db     *gorm.DB
	redis  *redis.Client
	l      zerolog.Logger
}

func NewAuthorsRepository(cfg *config.Config, db *gorm.DB, redis *redis.Client, log zerolog.Logger) AuthorsRepository {
	return &RepositoryAuthors{
		config: cfg,
		db:     db,
		redis:  redis,
		l:      log,
	}
}

func (r *RepositoryAuthors) CreateAuthor(author models.Author) (*models.Author, error) {
	if err := r.db.Create(&author).Error; err != nil {
		r.l.Error().
			Err(err).
			Str("module", "authors_repository").
			Str("function", "CreateAuthor").
			Msg("error creating author in DB")
		return nil, err
	}

	return &author, nil
}

func (r *RepositoryAuthors) GetAuthorByID(ctx context.Context, authorID int) (*models.Author, error) {
	var author models.Author

	cacheKey := "book:" + string(rune(authorID))

	cachedAuthor, err := r.redis.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		r.l.Info().
			Str("module", "authors_repository").
			Str("function", "GetAuthorByID").
			Int("author_id", authorID).
			Msg("fetching author by ID from DB")
		if err := r.db.First(&author, authorID).Error; err != nil {
			r.l.Error().
				Err(err).
				Str("module", "authors_repository").
				Str("function", "GetAuthorByID").
				Int("author_id", authorID).
				Msg("error getting author in DB")
			return nil, err
		}

		authorJSON, _ := json.Marshal(author)
		err = r.redis.Set(ctx, cacheKey, authorJSON, 10*time.Minute).Err()
		if err != nil {
			r.l.Error().
				Err(err).
				Str("module", "authors_repository").
				Str("function", "GetAuthorByID").
				Int("author_id", authorID).
				Msg("error setting author in redis")
			return nil, err
		}

		return &author, nil
	} else if err != nil {
		r.l.Error().
			Err(err).
			Str("module", "authors_repository").
			Str("function", "GetAuthorByID").
			Int("author_id", authorID).
			Msg("error searching author in redis")
		return nil, err
	}

	if err := json.Unmarshal([]byte(cachedAuthor), &author); err != nil {
		r.l.Error().
			Err(err).
			Str("module", "authors_repository").
			Str("function", "GetAuthorByID").
			Int("author_id", authorID).
			Msg("error unmarshaling author from redis")
		return nil, err
	}
	r.l.Info().
		Str("module", "authors_repository").
		Str("function", "GetAuthorByID").
		Int("author_id", authorID).
		Msg("fetching author by ID from Redis")
	return &author, nil
}

func (r *RepositoryAuthors) GetAuthors() ([]models.Author, error) {
	var authors []models.Author

	if err := r.db.Find(&authors).Error; err != nil {
		r.l.Error().
			Err(err).
			Str("module", "authors_repository").
			Str("function", "GetAuthors").
			Msg("error getting authors from DB")
		return nil, err
	}

	return authors, nil
}

func (r *RepositoryAuthors) UpdateAuthor(authorID int, author *models.Author) (*models.Author, error) {
	result := r.db.Model(&models.Author{}).Where("author_id = ?", authorID).Updates(author)
	if result.Error != nil {
		r.l.Error().
			Err(result.Error).
			Str("module", "authors_repository").
			Str("function", "UpdateAuthor").
			Int("author_id", authorID).
			Msg("error updating book in DB")
		return nil, result.Error // Retorna error si ocurrió algún problema
	}
	if result.RowsAffected == 0 {
		r.l.Error().
			Err(gorm.ErrRecordNotFound).
			Str("module", "authors_repository").
			Str("function", "UpdateAuthor").
			Int("author_id", authorID).
			Msg("error updating book in DB")
		return nil, gorm.ErrRecordNotFound // Retorna error si no se encontró el libro
	}

	author.AuthorID = authorID

	return author, nil
}

func (r *RepositoryAuthors) DeleteAuthor(authorID int) error {
	result := r.db.Delete(&models.Author{}, authorID)
	if result.Error != nil {
		r.l.Error().
			Err(result.Error).
			Str("module", "authors_repository").
			Str("function", "DeleteAuthor").
			Int("author_id", authorID).
			Msg("error deleting author from DB")
		return result.Error
	}
	if result.RowsAffected == 0 {
		r.l.Error().
			Err(gorm.ErrRecordNotFound).
			Str("module", "authors_repository").
			Str("function", "DeleteAuthor").
			Int("author_id", authorID).
			Msg("error deleting author from DB")
		return gorm.ErrRecordNotFound
	}
	return nil
}
