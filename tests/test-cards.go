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
func testCard() {
	// Create a valid card and print it
	myCard, err := card.NewCard("J", "♣")
	check(err)
	myCard.Print()
	// Create a card with invalid value
	badValue, err := card.NewCard("10", "♣")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Properly failed on invalid suit")
	} else {
		fmt.Print("Bad card got through: ")
		badValue.Print()
	}
}

func testShoe() {
	//Create Shoe, print the cards, shuffle it and reprint
	myShoe := shoe.NewShoe()
	for _, shoeCard := range myShoe.Pile {
		shoeCard.Print()
	}
	myShoe.Shuffle()
	fmt.Println("")
	for _, shoeCard := range myShoe.Pile {
		shoeCard.Print()
	}
}

func main() {
	// Test Card Package
	testCard()
	fmt.Println("")
	// Test Shoe Package
	testShoe()
	fmt.Println("")
	// Test Player Package
	newPlayer, err := player.NewPlayer("rhartman", 100)
	check(err)

	myShoe := shoe.NewShoe()
	card1, err1 := myShoe.DealCard()
	check(err1)
	card2, err2 := myShoe.DealCard()
	check(err2)

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
		* Realized that I need to be thinking of Shoes instead of shoes:
		** Probably need to rework the shoe module to be a shoe
		** Don't need to track individual shoes so construction would be similar
		** But having one shoe will make testing easier for now so that can wait */

	fmt.Println("Player Hands:")
	newPlayer.PrintHands()
}
