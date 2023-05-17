package deck

import "fmt"

func ExampleCard() {
  fmt.Println(Card{Rank: Ace, Suit: Spade})
  fmt.Println(Card{Rank: Nine, Suit: Heart})
  fmt.Println(Card{Rank: Queen, Suit: Club})

  // Output:
  // Ace of Spades
  // Nine of Hearts
  // Queen of Clubs
}
