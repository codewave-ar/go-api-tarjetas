package actions_test

import (
	"app-go/actions"
	"app-go/domain"
	"testing"
)

func TestGetCard_shouldBeSuccess(t *testing.T) {
	// Setup
	expectedCard := domain.Card{
		"1234123412341234",
		"Test test",
		"01/25",
		"VISA",
	}
	getter := actions.CardGetter{
		Cards: map[string]domain.Card{
			"1234123412341234": expectedCard,
		},
	}

	// Execute
	card, err := getter.Execute("1234123412341234")

	// Verify
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if card != expectedCard {
		t.Errorf("Distinct card")
	}
}

func TestGetCard_shouldBeNotFound(t *testing.T) {
	// Setup
	getter := actions.CardGetter{
		Cards: map[string]domain.Card{},
	}

	// Execute
	card, err := getter.Execute("1234123412341235")

	// Verify
	if err == nil {
		t.Fatalf("We expected an error")
	}
	if card.CardNumber != "" {
		t.Errorf("Unexpected card: %v", card)
	}
}
