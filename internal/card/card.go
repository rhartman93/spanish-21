package card

import (
	"errors"
)

var Suits = [4]string{"♣", "♦", "♥", "♠"}

var ValueMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	//"10": 10, No 10s in a Spanish 21 deck
	"J": 10,
	"Q": 10,
	"K": 10,
	"A": 11,
}

type Card struct {
	Value        int
	DisplayValue string
	Suit         string
}

func NewCard(value string, suit string) (Card, error) {
	thisCard := new(Card)
	// Check for valid value
	var ok bool
	thisCard.Value, ok = ValueMap[value]
	if !ok {
		return *thisCard, errors.New("invalid card value")
	}
	thisCard.DisplayValue = value
	// Check for valid suit
	for _, v := range Suits {
		if v == suit {
			thisCard.Suit = suit
			break
		}
	}
	if thisCard.Suit == "" {
		return *thisCard, errors.New("invalid card suit")
	}

	return *thisCard, nil
}

func (c Card) String() string {
	return "[" + c.DisplayValue + c.Suit + "]"
}
