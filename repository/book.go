package repository

import (
	"errors"

	"github.com/ainmtsn1999/orm-book-api-deploy/helper"
	"github.com/ainmtsn1999/orm-book-api-deploy/model"
)

// interface employee
type BookRepo interface {
	CreateBook(in model.GormBook) (res model.GormBook, err error)
	GetAllBook() (res []model.GormBook, err error)
	GetBookById(id int64) (res model.GormBook, err error)
	UpdateBook(in model.GormBook) (res model.GormBook, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) CreateBook(in model.GormBook) (res model.GormBook, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.GormBook, err error) {
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookById(id int64) (res model.GormBook, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.GormBook) (res model.GormBook, err error) {

	if (r.gorm.Model(&res).Where("id = ?", in.Id).Updates(model.GormBook{
		NameBook: in.NameBook,
		Author:   in.Author,
	}).Scan(&res).RowsAffected == 0) {
		return res, errors.New(helper.ErrNotFound)
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {

	if (r.gorm.Delete(&model.GormBook{}, id).RowsAffected == 0) {
		return errors.New(helper.ErrNotFound)
	}

	return nil
}
