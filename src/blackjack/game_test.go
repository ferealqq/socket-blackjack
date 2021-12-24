package blackjack

import (
	"testing"
)

func TestDealFor(t *testing.T){
	b := NewGame("pekka")
	c1 := b.DealFor(&b.Player)
	c2 := b.DealFor(&b.Player)
	if len(b.deck.cards ) != 52-2{
		t.Error("DealFor failed")
	}
	hand := b.Player.hand
	if len(hand) != 2{
		t.Error("DealFor failed because the Players hand is not of length 2")
	}
	if (hand[0].rank != c1.rank ||
			hand[1].rank != c2.rank){
		t.Error("DealFor failed because the Player does not have the correct values")
	}
}
