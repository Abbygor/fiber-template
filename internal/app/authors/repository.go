package authors

import (
	"fiber-template/internal/config"
	"fiber-template/internal/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthorsRepository interface {
	CreateAuthor(models.Author) (*models.Author, error)
	GetAuthorByID(int) (*models.Author, error)
	GetAuthors() ([]models.Author, error)
	UpdateAuthor(int, *models.Author) (*models.Author, error)
	DeleteAuthor(int) error
}

type RepositoryAuthors struct {
	config *config.Config
	db     *gorm.DB
	redis  *redis.Client
}

func NewAuthorsRepository(cfg *config.Config, db *gorm.DB, redis *redis.Client) AuthorsRepository {
	return &RepositoryAuthors{
		config: cfg,
		db:     db,
		redis:  redis,
	}
}

func (r *RepositoryAuthors) CreateAuthor(author models.Author) (*models.Author, error) {
	if err := r.db.Create(&author).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *RepositoryAuthors) GetAuthorByID(authorID int) (*models.Author, error) {
	var author models.Author

	if err := r.db.First(&author, authorID).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *RepositoryAuthors) GetAuthors() ([]models.Author, error) {
	var authors []models.Author

	if err := r.db.Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *RepositoryAuthors) UpdateAuthor(authorID int, author *models.Author) (*models.Author, error) {
	result := r.db.Model(&models.Author{}).Where("author_id = ?", authorID).Updates(author)
	if result.Error != nil {
		return nil, result.Error // Retorna error si ocurrió algún problema
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound // Retorna error si no se encontró el libro
	}

	author.AuthorID = authorID

	return author, nil
}

func (r *RepositoryAuthors) DeleteAuthor(authorID int) error {
	result := r.db.Delete(&models.Author{}, authorID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
