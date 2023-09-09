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

// Might need to use a pointer for shuffling since it manipulates the shoe
// should test
func (d Shoe) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Pile), func(i, j int) {
		d.Pile[i], d.Pile[j] = d.Pile[j], d.Pile[i]
	})
}

func NewShoe(numDecks int) (Shoe, error) {
	myShoe := new(Shoe)
	if numDecks < 0 {
		return *myShoe, errors.New("must have at least 1 deck")
	}
	for i := 1; i <= numDecks; i++ {
		for key := range card.ValueMap {
			for _, suit := range card.Suits {
				thisCard, err := card.NewCard(key, suit)
				if err != nil {
					fmt.Println("something went wrong adding cards to shoe")
					fmt.Println(err)
					os.Exit(-1)
				}
				myShoe.Pile = append(myShoe.Pile, thisCard)
			}
		}
	}
	return *myShoe, nil
}

func (d *Shoe) DealCard() (card.Card, error) {
	if len(d.Pile) == 0 {
		//If the deck is empty return a blank card and an error
		return *new(card.Card), errors.New("emptyShoe")
	}
	dealtCard := d.Pile[len(d.Pile)-1]
	d.Pile = d.Pile[:len(d.Pile)-1]

	return dealtCard, nil
}
