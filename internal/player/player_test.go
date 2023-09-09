package player

import (
	"reflect"
	"spanish21/internal/card"
	"spanish21/internal/shoe"
	"testing"
)

/*
	######################

Need to fix the test comparing the cards that they have,
might want to just remove the Shoe aspect, create the cards, and add them
*/

var expectedCards = []card.Card{
	{Value: 11, DisplayValue: "A", Suit: "♠"},
	{Value: 11, DisplayValue: "A", Suit: "♥"},
	{Value: 11, DisplayValue: "A", Suit: "♦"},
	{Value: 11, DisplayValue: "A", Suit: "♣"},
}

func TestCreatePlayer(t *testing.T) {
	newPlayer, errPlayer := NewPlayer("test-player-1", 100)
	if errPlayer != nil {
		t.Errorf("Error creating player: %s", errPlayer.Error())
	}

	if len(newPlayer.Hands) != 1 {
		t.Errorf("expected 1 hand for new player, instead have: %d", len(newPlayer.Hands))
	}
	myShoe, _ := shoe.NewShoe(1)
	card1, _ := myShoe.DealCard()
	card2, _ := myShoe.DealCard()

	newPlayer.Hands[0] = append(newPlayer.Hands[0], card1)
	newPlayer.Hands[0] = append(newPlayer.Hands[0], card2)

	card3, _ := myShoe.DealCard()
	card4, _ := myShoe.DealCard()
	newPlayer.AddHand([]card.Card{card3, card4})

	for x, hand := range newPlayer.Hands {
		for y, thisCard := range hand {
			if !reflect.DeepEqual(thisCard, expectedCards[x+y]) {
				t.Errorf("expected player to have card %s, but they had %s", expectedCards[x+y].String(), thisCard.String())
			}
		}
	}
}
