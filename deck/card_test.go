package deck

import "fmt"

func ExampleCard() {
  c := Card{Rank: Ace, Suit: Spade}
  fmt.Println(c.String())

  // Output:
  // Ace of Spades
}
