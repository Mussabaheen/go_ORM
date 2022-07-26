package bookstore

import "context"

type service struct {
	repository Repository
}

type Service interface {
	getBook(ctx context.Context) (getBookResponse, error)
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) getBook(ctx context.Context) (getBookResponse, error) {
	book, err := s.repository.getBook(ctx)
	if err != nil {
		return getBookResponse{}, err
	}
	return book, nil
}
