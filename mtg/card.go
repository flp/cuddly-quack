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

var CardRarityStringLookup = map[CardRarity]string{
	Common:    "Common",
	Uncommon:  "Uncommon",
	Rare:      "Rare",
	Mythic:    "Mythic",
	Special:   "Special",
	BasicLand: "BasicLand",
}

func (cr CardRarity) String() string {
	return CardRarityStringLookup[cr]
}

type Card struct {
	Name   string
	Rarity CardRarity
}

func (c Card) String() string {
	return fmt.Sprintf("NAME: %s, RARITY: %s", c.Name, c.Rarity)
}
