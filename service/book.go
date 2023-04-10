package service

import (
	"github.com/ainmtsn1999/orm-book-api-deploy/model"
)

// interface employee
type BookService interface {
	CreateBook(in model.GormBook) (res model.GormBook, err error)
	GetAllBook() (res []model.GormBook, err error)
	GetBookById(id int64) (res model.GormBook, err error)
	UpdateBook(in model.GormBook) (res model.GormBook, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.GormBook) (res model.GormBook, err error) {
	return s.repo.CreateBook(in)
}
func (s *Service) GetAllBook() (res []model.GormBook, err error) {
	return s.repo.GetAllBook()
}
func (s *Service) GetBookById(id int64) (res model.GormBook, err error) {
	return s.repo.GetBookById(id)
}
func (s *Service) UpdateBook(in model.GormBook) (res model.GormBook, err error) {
	return s.repo.UpdateBook(in)
}
func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
