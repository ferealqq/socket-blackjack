package blackjack

import (
	"fmt"
	"strconv"
)

type Card struct {
	rank int
	suit string
}

func (c Card) getRank() string {
	switch c.rank {
		case 1:
			return "Ace"
		case 11: 
			return "Jack"
		case 12:
			return "Queen"
		case 13:
			return "King"
		default:
			return strconv.Itoa(c.rank)
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.getRank(), c.suit)
}

func createPack() []Card {
	pack := make([]Card, 52)
	for i := 0; i < 52; i++ {
		number := i+1
		if((number % 13) == 0){
			pack[i].rank = 13
		}else {
			pack[i].rank = number % 13
		}
		if((number / 13) == 4){
			pack[i].suit = getSuit(3)
		}else{
			pack[i].suit = getSuit(number / 13)
		}
	}
	return pack
}

func getSuit(index int) string {
	switch index {
		case 0:
			return "Spades"
		case 1:
			return "Hearts"
		case 2:
			return "Diamonds"
		case 3:
			return "Clubs"
	}
	return "Unknown"
}
