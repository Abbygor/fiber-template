package authors

import (
	"context"
	"errors"
	"fiber-template/internal/models"
	"time"

	"github.com/rs/zerolog"
)

type AuthorsService interface {
	CreateAuthor(models.Author) (*models.Author, error)
	GetAuthorByID(context.Context, int) (*models.Author, error)
	GetAuthors() ([]models.Author, error)
	UpdateAuthor(int, *models.Author) (*models.Author, error)
	DeleteAuthor(int) error
}

type ServiceAuthors struct {
	authorRepository AuthorsRepository
	l                zerolog.Logger
}

func NewAuthorsService(repo AuthorsRepository, log zerolog.Logger) AuthorsService {
	return &ServiceAuthors{
		authorRepository: repo,
		l:                log,
	}
}

func (s *ServiceAuthors) CreateAuthor(author models.Author) (*models.Author, error) {
	if author.FirstName == "" || author.LastName == "" {
		s.l.Error().
			Str("module", "authors_service").
			Str("function", "CreateAuthor").
			Err(errors.New("full author's name required")).
			Msg("missing field: name")
		return nil, errors.New("full author's name required")
	}
	_, err := time.Parse("2006-01-02", author.BirthDate)
	if err != nil {
		s.l.Error().
			Str("module", "authors_service").
			Str("function", "CreateAuthor").
			Err(err).
			Msg("date format error")
		return nil, errors.New("date format error")
	}

	return s.authorRepository.CreateAuthor(author)

}

func (s *ServiceAuthors) GetAuthorByID(ctx context.Context, authorID int) (*models.Author, error) {
	author, err := s.authorRepository.GetAuthorByID(ctx, authorID)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "authors_service").
			Str("function", "GetAuthorByID").
			Int("author_id", authorID).
			Msg("error getting author from repository")
		return nil, err
	}
	return author, nil
}

func (s *ServiceAuthors) GetAuthors() ([]models.Author, error) {
	authors, err := s.authorRepository.GetAuthors()
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "authors_service").
			Str("function", "GetAuthors").
			Msg("error getting all authors from repository")
		return nil, err
	}
	if len(authors) == 0 {
		s.l.Error().
			Err(errors.New("no authors found")).
			Str("module", "authors_service").
			Str("function", "GetAuthors").
			Msg("no authors found")
		return nil, errors.New("no authors found")
	}

	return authors, nil
}

func (s *ServiceAuthors) UpdateAuthor(authorID int, author *models.Author) (*models.Author, error) {
	updatedAuthor, err := s.authorRepository.UpdateAuthor(authorID, author)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "authors_service").
			Str("function", "UpdateAuthor").
			Int("author_id", authorID).
			Msg("error updating author from repository")
		return nil, err
	}

	return updatedAuthor, nil
}

func (s *ServiceAuthors) DeleteAuthor(authorID int) error {
	err := s.authorRepository.DeleteAuthor(authorID)
	if err != nil {
		s.l.Error().
			Err(err).
			Str("module", "authors_service").
			Str("function", "DeleteAuthor").
			Int("author_id", authorID).
			Msg("error deleting author from repository")
		return err
	}
	return nil
}
