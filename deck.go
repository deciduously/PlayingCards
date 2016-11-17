//Package PlayingCards contains structures and functions implementing a standard deck of playing cards
package PlayingCards

import (
	"math/rand"
)

//Card represents a single card
type Card struct {
	value uint8
	suit  Suit
}

//Stack represents a stack of Cards
type Stack []Card

//Suit represents one of the four suits
type Suit string

//The four suits
const (
	Clubs    Suit = "clubs"
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Spades   Suit = "spades"
)

//NewDeck returns a new deck, consisting of a Stack of 52 Cards
func NewDeck() Stack {
	deck := make(Stack, 0)
	suits := []Suit{Clubs, Hearts, Diamonds, Spades}
	for _, i := range suits {
		for j := 1; j <= 13; j++ {
			deck = append(deck, Card{uint8(j), i})
		}
	}
	return deck
}

/*
//Deal deals n cards off of Stack d
func (d Stack) Deal (n uint) {

}
*/

//Shuffle shuffles Stack d
func (d Stack) Shuffle() {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
