package actions

import (
	"app-go/domain"
	"fmt"
)

type CardGetter struct {
	Cards map[string]domain.Card
}

func (cg *CardGetter) Execute(cardNumber string) (domain.Card, error) {
	card, isValueFound := cg.Cards[cardNumber]
	if !isValueFound {
		return domain.Card{}, fmt.Errorf("Card not found: %s", cardNumber)
	}
	return card, nil
}
