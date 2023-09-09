package shoe

import (
	"spanish21/internal/card"
	"testing"
)

// Create list of all legal suit/value pairs for deck sanity testing
type cardStub struct {
	displayValue, suit string
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
	//Generate card stubs to test deck against
	cardList := GenerateCardChecks()
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
		var count int
		for _, stub := range *cardList {
			count = 0
			for _, thisCard := range gotShoe.Pile {
				if thisCard.DisplayValue == stub.displayValue && thisCard.Suit == stub.suit {
					count++
				}
				if count == i {
					break
				}
			}
			if count != i {
				t.Errorf("expected %d instances of card [%s%s] in shoe of size %d but instead found %d", i, stub.displayValue, stub.suit, i, count)
			}
		}

	}
}

func TestDealingCard(t *testing.T) {
	//Can assume no error here since it's tested in a previous test
	newShoe, _ := NewShoe(1)
	wantCard1 := newShoe.Pile[len(newShoe.Pile)-1]
	wantCard2 := newShoe.Pile[len(newShoe.Pile)-2]
	t.Logf("%d cards in deck before first deal", len(newShoe.Pile))
	gotCard1, err1 := newShoe.DealCard()
	t.Logf("%d cards in deck after first deal", len(newShoe.Pile))
	gotCard2, err2 := newShoe.DealCard()

	if err1 != nil || err2 != nil {
		t.Errorf("error occurred trying to deal card")
	}
	if wantCard1 != gotCard1 || wantCard2 != gotCard2 {
		t.Errorf("expected dealt cards %s%s but got cards %s%s", wantCard1.String(), wantCard2.String(), gotCard1.String(), gotCard2.String())
	}
}
