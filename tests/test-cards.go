package main

import (
	"fmt"
	"os"
	"spanish21/internal/card"
	"spanish21/internal/player"
	"spanish21/internal/shoe"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	// Test Player Package
	newPlayer, errPlayer := player.NewPlayer("rhartman", 100)
	check(errPlayer)
	myShoe, errShoe1 := shoe.NewShoe(1)
	check(errShoe1)
	card1, errDeal1 := myShoe.DealCard()
	check(errDeal1)
	card2, errDeal2 := myShoe.DealCard()
	check(errDeal2)

	newPlayer.Hands[0] = append(newPlayer.Hands[0], card1)
	newPlayer.Hands[0] = append(newPlayer.Hands[0], card2)

	card3, err3 := myShoe.DealCard()
	check(err3)
	card4, err4 := myShoe.DealCard()
	check(err4)
	newPlayer.AddHand([]card.Card{card3, card4})
	/* ################################
		Current Problems:
		* Hand 0 and 1 are the same
	    * Having to error check every deal of the cards is annoying
		* Convert to proper Go unit testing format
	*/

	fmt.Println("Player Hands:")
	newPlayer.PrintHands()
}
