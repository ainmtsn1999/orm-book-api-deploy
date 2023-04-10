package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type GormBook struct {
	Id        int       `json:"id" gorm:"primaryKey;type:serial"`
	NameBook  string    `json:"name_book" gorm:"type:varchar(200)"`
	Author    string    `json:"author" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g GormBook) Validation() error { // custom validation
	return validation.ValidateStruct(&g,
		validation.Field(&g.NameBook, validation.Length(5, 0), validation.Required),
		validation.Field(&g.Author, validation.Length(5, 0), validation.Required))
}
