//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
  Spade Suit = iota
  Diamond
  Club
  Heart
  Joker
)

type Rank uint8

const (
  _ Rank = iota
  Ace
  Two
  Three
  Four
  Five
  Six
  Seven
  Eight
  Nine
  Ten
  Jack
  Queen
  King
)

type Card struct {
  Suit
  Rank
}

func (c Card) String() string {
  if c.Suit == Joker {
    return c.Suit.String()
  }
  return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
  var cards []Card
  for i := 1; i <= 13; i++ {
    for j := 0; j <= 3; j++ {
      cards = append(cards, Card{Suit(j), Rank(i)})
    }
  }
  return cards
}

