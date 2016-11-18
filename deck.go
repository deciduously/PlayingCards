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
	Value uint8
	Suit  Suit
}

//Deck represents a Deck of Cards as a LIFO stack
type Deck struct {
  Cards []Card
}

//Suit represents one of the four suits
type Suit uint8

//The four suits
const (
	Clubs    Suit = 0x63
	Hearts   Suit = 0x68
	Diamonds Suit = 0x64
	Spades   Suit = 0x73
)

//Draw draws n cards off of Deck d, returns created Deck Deck, and the now smaller Deck f
//If n > len(d) or n < 1, it simply returns an empty Deck and Deck d untouched, as well as an error
func (d *Deck) Draw(n int) (ret Deck, e error) {
	ret = Deck{}
	if n > d.Len() || n < 1 {
		e = fmt.Errorf("Could not draw %v from Deck of size %v", n, d.Len())
		return ret, e
	}
	for i := 0; i < n; i++ {
		ret.Push(d.Pop())
	}
	return ret, nil
}

//Empty return true if Deck d is empty, otherwise false
func (d Deck) Empty() bool {
	return d.Len() == 0
}

//Len returns to size of the Deck
func (d *Deck) Len() int {
  return len(d.Cards)
}


//NewDeck returns a new deck, consisting of a Deck of 52 Cards
func NewDeck() Deck {
	d := Deck{}
	suits := []Suit{Clubs, Hearts, Diamonds, Spades}
	for _, i := range suits {
		for j := 1; j <= 13; j++ {
			d.Push(Card{uint8(j), i})
		}
	}
	return d
}

//Peek returns the top card of Deck d
func (d Deck) Peek() Card {
	return d.Cards[d.Len() -1]
}

//Pop pops the top card from the Deck
func (d *Deck) Pop() Card {
  c := d.Cards[len(d.Cards)-1]
  d.Cards = d.Cards[:len(d.Cards)-1]
  return c
}

//Push pushes a Card to a Deck
func (d *Deck) Push(c Card) {
  d.Cards = append(d.Cards, c)
}

//Readable returns a string containing the Card info for output
func (c Card) Readable() string {
	val := ""
  suit := ""
	switch c.Value {
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
		val = strconv.Itoa(int(c.Value))
	}

  switch c.Suit {
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

//Shuffle shuffles Deck d
func (d *Deck) Shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d.Cards {
		j := r.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
