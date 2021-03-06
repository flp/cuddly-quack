package mtg

import (
	"fmt"
	"math/rand"
)

type Pack struct {
	Cards []*Card
}

func NewPack() *Pack {
	return &Pack{
		Cards: []*Card{},
	}
}

func (p Pack) String() string {
	str := ""
	for _, card := range p.Cards {
		str += fmt.Sprintf("%s\n", card)
	}
	return str
}

type PackGenerator interface {
	GeneratePack() Pack
}

type ModernPackGenerator struct {
}

// ChooseNCards chooses n cards at random from an array of cards.
func ChooseNCards(cards []*Card, n int) []*Card {
	if n > len(cards) {
		// TODO(flp): better error handling
		return nil
	}

	shuffle := func(a []*Card) {
		for i := range a {
			j := rand.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}
	}

	var result []*Card
	var cardsCopy = make([]*Card, len(cards))
	copy(cardsCopy, cards)
	for i := 0; i < n; i++ {
		shuffle(cardsCopy)
		result = append(result, cardsCopy[0])
		cardsCopy = cardsCopy[1:]
	}

	return result
}

// ModernPackGenerator's GeneratePack creates a pack with 14 cards: 1 random rare, 3 random
// uncommons, and 10 random commons.
func (mpg ModernPackGenerator) GeneratePack(s SetName) (Pack, error) {
	// TODO(flp): foils
	p := *NewPack()
	raritiesToCards, err := GetCardsForSet(s)
	if err != nil {
		return p, err
	}

	// Get rare (possibly a mythic rare)
	n := rand.Intn(8)
	if n == 7 {
		// Choose a mythic rare instead of a normal rare
		mythics := raritiesToCards[Mythic]
		p.Cards = append(p.Cards, ChooseNCards(mythics, 1)...)
	} else {
		// Choose a normal rare
		rares := raritiesToCards[Rare]
		p.Cards = append(p.Cards, ChooseNCards(rares, 1)...)
	}

	// Get uncommons
	uncommons := raritiesToCards[Uncommon]
	p.Cards = append(p.Cards, ChooseNCards(uncommons, 3)...)

	// Get commons
	commons := raritiesToCards[Common]
	p.Cards = append(p.Cards, ChooseNCards(commons, 10)...)

	return p, nil
}
