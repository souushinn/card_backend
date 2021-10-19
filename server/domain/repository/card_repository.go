package repository

import (
	"context"

	"github.com/souushinn/cardGo/domain/model"
)

type CardRepository interface {
	ListCards(ctx context.Context, params ListCardsParams) ([]*model.Card, error)
}

type ListCardsParams struct {
	Limit int
}
