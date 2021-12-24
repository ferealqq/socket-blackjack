package main

import (
	"blackjack-game/src/blackjack"
)

func main() {	
	game := blackjack.NewGame("pekka")
	win := game.Play(&game.Player)
	if !win {
		game.PlayDealer()
	}
}