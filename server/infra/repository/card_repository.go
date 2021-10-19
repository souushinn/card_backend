package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/souushinn/cardGo/domain/domainerror"
	"github.com/souushinn/cardGo/domain/model"
	"github.com/souushinn/cardGo/domain/repository"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) repository.CardRepository {
	return &CardRepository{
		db: db,
	}
}

func (c *CardRepository) ListCards(ctx context.Context, params repository.ListCardsParams) ([]*model.Card, error) {
	db := c.db.Order("id")
	var Cards []*model.Card
	if err := db.Find(&Cards).Error; err != nil {
		return nil, domainerror.NewInternalServerError(err.Error(), err)
	}
	return Cards, nil
}
