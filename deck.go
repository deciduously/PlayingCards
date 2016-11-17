//Package PlayingCards contains structures and functions implementing a standard deck of playing cards
package PlayingCards

//Card represents a single card
type Card struct {
	value uint8
	suit  Suit
}

//Deck represents a deck of Cards
type Deck []Card

//Suit represents one of the four suits
type Suit string

//The four suits
const (
	Clubs    Suit = "clubs"
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Spades   Suit = "spades"
)

//NewDeck returns a new Deck of Cards,
func NewDeck() Deck {
	deck := make(Deck, 0)
	suits := []Suit{Clubs, Hearts, Diamonds, Spades}
	for _, i := range suits {
		for j := 1; j <= 13; j++ {
			deck = append(deck, Card{uint8(j), i})
		}
	}
	return deck
}

/*
//ShuffleDeck shuffles and returns Deck
func ShuffleDeck(d *Deck) {

}
*/
