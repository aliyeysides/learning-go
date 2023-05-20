package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Heart})
	fmt.Println(Card{Rank: Queen, Suit: Club})

	// Output:
	// Ace of Spades
	// Nine of Hearts
	// Queen of Clubs
}

func ExampleNew() {
	fmt.Println(len(New()))
	fmt.Println(New(Sort(Less))[0])

	// Output:
	// 52
	// Ace of Spades
}

// Same test as above but done with testing package. This approach has the benefit of throwing specific error messages.
func TestNew(t *testing.T) {
	cards := New()
	// 52 cards in a deck
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	// First card should be Ace of Spades
	if cards[0] != (Card{Suit: Spade, Rank: Ace}) {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	// First card should be Ace of Spades
	if cards[0] != (Card{Suit: Spade, Rank: Ace}) {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	cards := New(DefaultSort, Shuffle)
	// First card should not be Ace of Spades
	if cards[0] == (Card{Suit: Spade, Rank: Ace}) && cards[1] == (Card{Suit: Spade, Rank: Two }) {
		t.Error("Expected first and second card to not be Ace of Spades and Two of Spades. Received Both")
	}
}
