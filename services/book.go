package services

import (
	"gomvc/models"
	"gomvc/repos"
)

type BookService interface {
	GetAll() []models.Book
	GetByID(id int64) models.Book
	GetByName(name string) models.Book
}

type bookService struct {
	repo repos.BookRepository
}

func NewBookService(repo repos.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) GetAll() []models.Book {
	return s.repo.Select("select * from books")
}

func (s *bookService) GetByID(id int64) models.Book {
	return s.repo.SelectById("select * from books where ID=?", id)
}

func (s *bookService) GetByName(name string) models.Book {
	return s.repo.SelectByName("select * from books where title=?", name)
}
