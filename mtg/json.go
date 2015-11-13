package mtg

import (
	"encoding/json"
)

// RaritiesToCards maps card rarities to lists of Cards. This map is used for generating packs.
// This map holds all cards that we can possibly generate. This should be regarded as a temporary
// solution, since we probably don't want to keep tens of thousands of cards in memory.
var RaritiesToCards = make(map[CardRarity][]*Card)

// ReadCardData encodes JSON data from a byte array from a single set into a Set struct.
// The Cards from the Set are stored in the RaritiesToCards global var.
func ReadCardData(contents []byte) {
	var s Set
	_ = json.Unmarshal(contents, &s)
	// TODO(flp): better error handling

	getCardRarity := func(c *Card) CardRarity {
		switch {
		case c.Rarity == "Common":
			return Common
		case c.Rarity == "Uncommon":
			return Uncommon
		case c.Rarity == "Rare":
			return Rare
		case c.Rarity == "Mythic Rare":
			return Mythic
		}
		return BasicLand
	}

	for _, card := range s.Cards {
		rarity := getCardRarity(card)
		card_arr := RaritiesToCards[rarity]
		RaritiesToCards[rarity] = append(card_arr, card)
	}
}
