package authors

import (
	"errors"
	"fiber-template/internal/models"
	"time"
)

type AuthorsService interface {
	CreateAuthor(models.Author) (*models.Author, error)
	GetAuthorByID(int) (*models.Author, error)
	GetAuthors() ([]models.Author, error)
	UpdateAuthor(int, *models.Author) (*models.Author, error)
	DeleteAuthor(int) error
}

type ServiceAuthors struct {
	authorRepository AuthorsRepository
}

func NewAuthorsService(repo AuthorsRepository) AuthorsService {
	return &ServiceAuthors{
		authorRepository: repo,
	}
}

func (s *ServiceAuthors) CreateAuthor(author models.Author) (*models.Author, error) {
	if author.FirstName == "" || author.LastName == "" {
		return nil, errors.New("full author's name required")
	}
	_, err := time.Parse("2006-01-02", author.BirthDate)
	if err != nil {
		return nil, errors.New("date format error")
	}

	return s.authorRepository.CreateAuthor(author)

}

func (s *ServiceAuthors) GetAuthorByID(authorID int) (*models.Author, error) {
	author, err := s.authorRepository.GetAuthorByID(authorID)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *ServiceAuthors) GetAuthors() ([]models.Author, error) {
	authors, err := s.authorRepository.GetAuthors()
	if err != nil {
		return nil, err
	}
	if len(authors) == 0 {
		return nil, errors.New("no authors found")
	}

	return authors, nil
}

func (s *ServiceAuthors) UpdateAuthor(authorID int, author *models.Author) (*models.Author, error) {
	updatedAuthor, err := s.authorRepository.UpdateAuthor(authorID, author)
	if err != nil {
		return nil, err
	}

	return updatedAuthor, nil
}

func (s *ServiceAuthors) DeleteAuthor(authorID int) error {
	err := s.authorRepository.DeleteAuthor(authorID)
	if err != nil {
		return err
	}
	return nil
}
