package card

import (
	"testing"
)

func TestGoodCard(t *testing.T) {
	// Create a valid card and print it
	gotCard, gotErr := NewCard("J", "♣")
	wantDisplayValue := "J"
	wantValue := 10
	wantSuit := "♣"
	if gotErr != nil {
		t.Errorf("got error trying to create card: %s", gotErr.Error())
	}
	if gotCard.DisplayValue != wantDisplayValue || gotCard.Value != wantValue || gotCard.Suit != wantSuit {
		t.Errorf("expected card [%s], got [%s]", wantDisplayValue+wantSuit, gotCard.DisplayValue+gotCard.Suit)
	}

}

func TestBadCard(t *testing.T) {
	// Create a card with invalid value
	_, err := NewCard("10", "♣")
	if err == nil {
		t.Errorf("creating invalid card should have returned error")
	}
}
