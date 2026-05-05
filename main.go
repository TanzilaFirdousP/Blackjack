package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	suit string
	rank string
}

type deck struct {
	cards []card
}

type hand struct {
	inhand []card
	score  int
}

// initializing deck
func (d *deck) deckinit() {
	suits := []string{"♥", "♦", "♣", "♠"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	d.cards = []card{}
	for _, s := range suits {
		for _, r := range ranks {
			d.cards = append(d.cards, card{s, r})
		}
	}
}

// shuffling deck
func (d *deck) shuffle() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	fmt.Println("Shuffled")
}

// to stack and shuffle before every round
func stack_shuffle() *deck {

	d := new(deck)
	d.deckinit()
	d.shuffle()
	return d

}

// drawing cards from the deck
func (d *deck) drawcard() card {
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c
}

// clearing the cards inhand of the players before every round

func (h *hand) total() int {
	total := 0
	aces := 0

	for _, c := range h.inhand {
		switch c.rank {
		case "J", "Q", "K":
			total += 10
		case "A":
			total += 11
			aces++
		default:
			var v int
			fmt.Sscanf(c.rank, "%d", &v)
			total += v
		}
	}

	for total > 21 && aces > 0 {
		total -= 10
		aces--
	}
	return total
}

// getter functions
type blackj interface {
	getcards() []card
	addcard(card) //addcard
	clear()       //clear inhand
	addscore()
	getscores() int
}

type player struct {
	hand
}

type dealer struct {
	hand
}

func (p *player) getcards() []card {
	return p.inhand
}
func (d *dealer) getcards() []card {
	return d.inhand
}

func (p *player) addcard(dcard card) {
	p.inhand = append(p.inhand, dcard)
}

func (d *dealer) addcard(dcard card) {
	d.inhand = append(d.inhand, dcard)
}

func addhand(a blackj, dcard card) {
	a.addcard(dcard)
}

func (p *player) addscore() {
	p.score++
}
func (d *dealer) addscore() {
	d.score++
}
func (p *player) getscores() int {
	return p.score
}
func (d *dealer) getscores() int {
	return d.score
}

func (p *player) clear() {
	p.inhand = make([]card, 0)
}

func (d *dealer) clear() {
	d.inhand = make([]card, 0)
}

func chand(cl blackj) {
	cl.clear()
}

// main function
func main() {
	p := new(player)
	d := new(dealer)
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	var busted int
	for {
		chand(p)
		chand(d)

		dt := 0
		pt := 0
		fmt.Println("1. Start Game") //to get choice
		fmt.Println("2. Exit")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			return
		}

		if choice != 1 && choice != 2 {
			fmt.Println("Invalid move. Please enter 1 or 2.")
			continue
		}
		//exit condition
		if choice == 2 {
			fmt.Println("Final Score -> Player:", p.getscores(), "Dealer:", d.getscores())
			if p.getscores() > d.getscores() {
				fmt.Println("YOU WON!!")
			} else if p.getscores() == d.getscores() {
				fmt.Println("DRAW")
			} else {
				fmt.Println("YOU LOSE :(")
			}
			break
		}

		// Reset

		deck := stack_shuffle()

		//cards are dealt
		card1 := deck.drawcard()
		addhand(p, card1)
		fmt.Println("Player dealt:", card1.rank, "of", card1.suit)

		card2 := deck.drawcard()
		addhand(d, card2)
		fmt.Println("Dealer dealt: [Hidden]")

		card3 := deck.drawcard()
		addhand(p, card3)
		fmt.Println("Player dealt:", card3.rank, "of", card3.suit)

		card4 := deck.drawcard()
		addhand(d, card4)
		fmt.Println("Dealer dealt:", card4.rank, "of", card4.suit)

		//blackjack
		if p.total() == 21 {
			fmt.Println("Blackjack! YOU WIN!")
			p.addscore()

			continue
		}
		busted = 0
		// Player turn
		for {
			fmt.Println("Total:", p.total())
			fmt.Println("1. Hit  2. Stand")

			var move int
			_, err := fmt.Scanln(&move)
			if err != nil {
				fmt.Println("Invalid input. Please enter 1 or 2.")
				continue
			}

			if move != 1 && move != 2 {
				fmt.Println("Invalid move. Please enter 1 or 2.")
				continue
			}

			if move == 1 {
				card := deck.drawcard()
				addhand(p, card)
				fmt.Println("Drew:", card.rank, card.suit)

				if p.total() > 21 {
					fmt.Println("Bust! Dealer wins!")
					d.addscore()
					busted++
					break
				}
				if p.total() == 21 {
					fmt.Println("21! YOU WIN!")
					p.addscore()
					busted++
					break
				}
			} else {
				break
			}
		}
		if busted == 0 {
			// Dealer turn if player not busted
			fmt.Println("Dealer reveals:")
			for _, c := range d.getcards() {
				fmt.Println(c.rank, c.suit)
			}

			for d.total() < 17 {
				card := deck.drawcard()
				addhand(d, card)
				fmt.Println("Dealer draws:", card.rank, card.suit)
			}

			dt = d.total()
			pt = p.total()

			fmt.Println("Dealer:", dt, "Player:", pt)

			//checking to declare winner
			if dt > 21 || (pt > dt) {
				fmt.Println("Player wins!")
				p.addscore()
			} else if dt > pt {
				fmt.Println("Dealer wins!")
				d.addscore()
			} else {
				fmt.Println("Draw!")
			}
		}
	}
}
