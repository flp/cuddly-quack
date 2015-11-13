package mtg

import (
	"fmt"
)

type CardRarity int

const (
	Common    CardRarity = 0
	Uncommon  CardRarity = 1
	Rare      CardRarity = 2
	Mythic    CardRarity = 3
	Special   CardRarity = 4
	BasicLand CardRarity = 5
)

type Card struct {
	Name   string "name"
	Rarity string "rarity"
}

type Set struct {
	Name  string  "name"
	Cards []*Card "cards"
}

func (c Card) String() string {
	return fmt.Sprintf("NAME: %s, RARITY: %s", c.Name, c.Rarity)
}
