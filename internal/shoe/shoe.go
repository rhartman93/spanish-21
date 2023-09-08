package shoe

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"spanish21/internal/card"
	"time"
)

type Shoe struct {
	Pile []card.Card
}

// Might need to use a pointer for shuffling since it manipulates the deck
// should test
func (d Shoe) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Pile), func(i, j int) {
		d.Pile[i], d.Pile[j] = d.Pile[j], d.Pile[i]
	})
}

func NewShoe() Shoe {
	myShoe := new(Shoe)
	for key := range card.ValueMap {
		for _, suit := range card.Suits {
			thisCard, err := card.NewCard(key, suit)
			if err != nil {
				fmt.Println("something went wrong adding cards to deck")
				fmt.Println("err")
				os.Exit(-1)
			}
			myShoe.Pile = append(myShoe.Pile, thisCard)
		}
	}
	return *myShoe
}

func (d *Shoe) DealCard() (card.Card, error) {
	if len(d.Pile) == 0 {
		//If the deck is empty return a blank card and an error
		return *new(card.Card), errors.New("emptyShoe")
	}
	dealtCard := d.Pile[len(d.Pile)-1]
	afterShoe := d.Pile[:len(d.Pile)-1]
	copy(d.Pile, afterShoe)

	return dealtCard, nil
}
