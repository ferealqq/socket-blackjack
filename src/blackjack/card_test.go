package blackjack

import (
	"testing"
)

func TestCreatePack(t *testing.T){
	c := createPack()
	if (c[0].getRank() != "Ace" || 
			c[0].rank != 1||
			c[0].suit != "Spades" ||
			c[51].getRank() != "King" ||
			c[51].rank != 13 ||
			c[51].suit != "Clubs"){
		t.Error("createPack failed")
	}
}