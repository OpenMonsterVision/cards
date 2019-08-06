package main

import "fmt"
import "math/rand"
import "strconv"
import "time"

type Card struct {
	value int    `json:"value"`
	suit  string `json:"suit"`
	color string `json:"color"`
	rank  string
}

type Deck struct {
	amount int
	cards  []Card `json:"deck"`
}

type Suit struct {
	name  string
	color string
}

func (d *Deck) generate() {
	c := Card{}
	for i := range suits {
		values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		for n := range values {
			switch values[n] {
			case 11:
				c.rank = "J"
			case 12:
				c.rank = "Q"
			case 13:
				c.rank = "K"
			case 14:
				c.rank = "A"
			default:
				c.rank = strconv.Itoa(values[n])
			}
			c.value = values[n]
			c.suit = suits[i].name
			c.color = suits[i].color
			d.cards = append(d.cards, c)
		}
	}
	d.shuffle()
}

func (d *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
func (d *Deck) split() (d1, d2 Deck) {
	d1.cards = d.cards[:26]
	d2.cards = d.cards[26:]
	return

}

func challenge() {
	if 1 > 2 {

	} else if 1 < 2 {
	} else {

	}
}

var suits = []Suit{
	{name: "Spades", color: "black"},
	{name: "Hearts", color: "red"},
	{name: "Diamonds", color: "red"},
	{name: "Clubs", color: "black"}}

func main() {
	fmt.Println(suits)
	var D Deck
	D.generate()
	p1, p2 := D.split()
}
