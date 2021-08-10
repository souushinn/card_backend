package repository

import "github.com/jinzhu/gorm"

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) repository.CardRepository {
	return &CardRepository{
		db: db,
	}
}
