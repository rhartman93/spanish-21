package player

import (
	"errors"
	"fmt"
	"spanish21/internal/card"
)

type Player struct {
	DisplayName string
	Hands       [][]card.Card
	Money       int
}

func NewPlayer(name string, money int) (Player, error) {
	thisPlayer := new(Player)
	// Check if money is valid before initializing anything else
	if money < 0 {
		return *thisPlayer, errors.New("player can't have negative money")
	}
	thisPlayer.DisplayName = name
	thisPlayer.Money = money
	thisPlayer.Hands = make([][]card.Card, 1)

	return *thisPlayer, nil
}

func (p *Player) AddHand(cards []card.Card) {
	// Get num of Hands before adding, this will translate to the index value of the new hand after adding (slice starts with 0)
	newHandNum := len(p.Hands)
	p.Hands = append(p.Hands, make([]card.Card, len(cards)))
	copy(p.Hands[newHandNum], cards)
}

func (p Player) PrintHands() {
	for _, hand := range p.Hands {
		for _, thisCard := range hand {
			fmt.Print(thisCard.String())
		}
	}
}
