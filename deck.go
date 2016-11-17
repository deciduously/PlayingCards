//Package playingcards implements a standard deck of playing cards
package playingcards

import (
	"fmt"
	"math/rand"
	"time"
  "strconv"
)

//Card represents a single card
type Card struct {
	value uint8
	suit  Suit
}

//Stack represents a stack of Cards
type Stack []Card

//Suit represents one of the four suits
type Suit uint8

//The four suits
const (
	Clubs    Suit = 0x63
	Hearts   Suit = 0x68
	Diamonds Suit = 0x64
	Spades   Suit = 0x73
)

//Draw draws n cards off of Stack d, returns created Stack stack, and the now smaller Stack f
//If n > len(d) or n < 1, it simply returns an empty Stack and Stack d untouched, as well as an error
func (d Stack) Draw(n int) (stack Stack, f Stack, e error) {
	stack = make(Stack, 0)
	if n > len(d) || n < 1 {
		e = fmt.Errorf("Could not draw %v from stack of size %v", n, len(d))
		return stack, d, e
	}
	for i := 0; i < n; i++ {
		stack = append(stack, d[i])
	}
	f = d[n:]
	return stack, f, nil
}

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

//Readable returns a string containing the Card info for output
func (c Card) Readable() (s string) {
	val := ""
  suit := ""
	switch c.value {
	case 0xa:
		val = "10"
	case 0xb:
		val = "Jack"
	case 0xc:
		val = "Queen"
	case 0xd:
		val = "King"
	case 0x1:
		val = "Ace"
	default:
		val = strconv.Itoa(int(c.value))
	}

  switch c.suit {
  case Clubs:
    suit = "Clubs"
    case Hearts:
    suit = "Hearts"
    case Diamonds:
    suit = "Diamonds"
    case Spades:
    suit = "Spades"
  }
  return val + " of " + suit
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
