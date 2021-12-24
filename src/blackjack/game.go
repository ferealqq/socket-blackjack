package blackjack

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func NewGame(playerName string) Blackjack{
	blackjack := Blackjack{
		deck: Deck{
			cards: createPack(),
		},
		Player: Player{
			name: playerName,
			hand: []Card{},
			score: 0,
			dealer: false,
		},
		dealer: Player{
			name: "Dealer",
			hand: []Card{},
			score: 0,
			dealer: true,
		},
	}

	blackjack.deck.shuffle()

	return blackjack
}

type Blackjack struct {
	deck Deck
	Player Player
	dealer Player
}

func (b *Blackjack) DealFor(p *Player) Card {
	c := b.deck.deal()
	p.addCard(c)
	return c
}

func (b *Blackjack) playerWon(){
	fmt.Println("You won the game!")
}

func (b *Blackjack) gameover() {
	fmt.Println("You lost the game!")
}

func (b *Blackjack) Play(player *Player) bool {
	reader := bufio.NewReader(os.Stdin)
	for{
		card := b.DealFor(player)
		fmt.Println("You've got an a card: " + card.String())
		fmt.Println("Current: " + b.Player.String())		
		if b.Player.hasBusted() {
			b.gameover()
			return false
		}
		if b.Player.isBlackjack() {
			b.playerWon();
			return true
		}
		fmt.Println("Do you want to hit? (Y/n)")
		text, _ := reader.ReadString('\n')
		if text == "n\n" {
			return false
		}else{
			continue
		}
	}
}

func (b *Blackjack) PlayDealer() {
	for {
		card := b.DealFor(&b.dealer)
		fmt.Println("Dealer got a card: " + card.String())
		fmt.Println("Dealer: " + b.dealer.String())
		if(b.dealer.isBlackjack()){
			b.gameover()
			return
		}else if b.dealer.hasBusted(){
			b.playerWon()
			break
		}else if b.dealer.score > b.Player.score {
			b.gameover()
			break
		}else{
			continue
		}
	}
}


type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d.cards {
		// length of cards
		j := rand.Intn(len(d.cards) - 0) + 0
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) deal() Card {
	c := d.cards[0]
	d.cards = d.cards[1:]
	d.shuffle()
	return c
}

type Player struct {
	name string
	hand []Card
	score int
	dealer bool
}

func (p *Player) addCard(card Card) {
	p.hand = append(p.hand, card)
	p.score += card.rank
}

func (p *Player) hasBusted() bool {
	return p.score > 21
}

// is blackjack
func (p *Player) isBlackjack() bool {
	return p.score == 21
}

// player to string
func (p *Player) String() string {
	return p.name + ": " + strconv.Itoa(p.score)
}