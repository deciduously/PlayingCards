//Package PlayingCards contains structures and functions implementing a standard deck of playing cards
package PlayingCards

import (
	"math/rand"
	"time"
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
	Clubs    Suit = "c"
	Hearts   Suit = "h"
	Diamonds Suit = "d"
	Spades   Suit = "s"
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

//Draw draws n cards off of Stack d, returns created Stack stack, and the smaller Stack f
func (d Stack) Draw(n int) (stack Stack, f Stack) {
	stack = make(Stack, 0)
	for i := 0; i < n; i++ {
		stack = append(stack, d[i])
	}
	f = d[n:]
	return stack, f
}

//Shuffle shuffles Stack d
func (d Stack) Shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		j := r.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
