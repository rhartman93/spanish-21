package shoe

import (
	"reflect"
	"spanish21/internal/card"
	"testing"
)

// Create list of all legal suit/value pairs for deck sanity testing
type cardStub struct {
	displayValue, suit string
}

func confirmShoeContents(numDecks int, testShoe Shoe, t *testing.T) {
	// Search shoe for x instances of card where x = numDecks
	// might be able to replace this with reflect.DeepEqual()
	cardList := GenerateCardChecks()
	for _, stub := range *cardList {
		count := 0
		for _, thisCard := range testShoe.Pile {
			if thisCard.DisplayValue == stub.displayValue && thisCard.Suit == stub.suit {
				count++
			}
			if count == numDecks {
				break
			}
		}
		if count != numDecks {
			t.Errorf("expected %d instances of card [%s%s] in shoe of size %d but instead found %d", numDecks, stub.displayValue, stub.suit, numDecks, count)
		}
	}

}

func GenerateCardChecks() *[]cardStub {
	var cs []cardStub
	for key := range card.ValueMap {
		for _, suit := range card.Suits {
			var thisCS cardStub
			thisCS.displayValue = key
			thisCS.suit = suit
			cs = append(cs, thisCS)
		}
	}
	return &cs
}
func TestCreateValidShoe(t *testing.T) {
	//Try creating shoes with 1-8 decks
	for i := 1; i < 8; i++ {
		gotShoe, err := NewShoe(i)
		//Check there wasn't an error making the Shoe
		if err != nil {
			t.Errorf("encountered an error when creating shoe: %s", err.Error())
		}
		//Check the proper number of cards are there
		if len(gotShoe.Pile) != i*48 {
			t.Errorf("expected shoe with %d cards but it contains %d", i*48, len(gotShoe.Pile))
		}
		// Search shoe for x instances of card where x = numDecks
		confirmShoeContents(i, gotShoe, t)
	}
}

func TestDealingCards(t *testing.T) {
	/* Overall test:
	deal two cards and make sure they're what you expect
	also make sure they're actually removed from the shoe
	finally deal the rest of the shoe, and make sure you don't
	get an error till you try to deal when the shoe's empty
	*/
	//Can assume no error here since it's tested in a previous test
	newShoe, _ := NewShoe(1)
	wantCard1 := newShoe.Pile[len(newShoe.Pile)-1]
	wantCard2 := newShoe.Pile[len(newShoe.Pile)-2]
	gotCard1, err1 := newShoe.DealCard()
	gotCard2, err2 := newShoe.DealCard()

	if err1 != nil || err2 != nil {
		t.Errorf("error occurred trying to deal card")
	}
	if wantCard1 != gotCard1 || wantCard2 != gotCard2 {
		t.Errorf("expected dealt cards %s%s but got cards %s%s", wantCard1.String(), wantCard2.String(), gotCard1.String(), gotCard2.String())
	}
	//Next Step, deal rest of deck + 1, checking for errors, expect error only on final deal when deck is empty
	for i := 0; i <= len(newShoe.Pile)+1; i++ {
		_, err := newShoe.DealCard()
		if i >= len(newShoe.Pile) && err != nil {
			t.Errorf("There was an error dealing cards when cards still left in shoe: %s", err.Error())
		} else if i <= 0 && len(newShoe.Pile) == 0 && err == nil {
			t.Errorf("didn't return error despite trying to deal from empty deck")
		}
	}
}

func TestShuffling(t *testing.T) {
	//Create a shoe, keep a copy, run shuffle, ensure the new one isn't in the same order and still has all cards
	shuffleShoe, _ := NewShoe(1)
	initialPile := make([]card.Card, len(shuffleShoe.Pile))
	copy(initialPile, shuffleShoe.Pile)

	shuffleShoe.Shuffle()
	if len(shuffleShoe.Pile) != len(initialPile) {
		t.Errorf("shuffling changed amount of cards in deck, initial")
	}
	confirmShoeContents(1, shuffleShoe, t)

	if reflect.DeepEqual(initialPile, shuffleShoe.Pile) {
		t.Errorf("deck order wasn't changed after shuffling")
	}
}
