package mtg

import (
	"encoding/json"

	"github.com/flp/cuddly-quack/assets"
)

type mtgjsonCard struct {
	Name   string "name"
	Rarity string "rarity"
}

type mtgjsonSet struct {
	Name  string         "name"
	Cards []*mtgjsonCard "cards"
}

// SetNamesToJsonResources maps SetNames to their json source files.
var SetNamesToJsonResources = map[SetName]string{
	BattleForZendikar: "resources/BFZ.json",
}

// SetsToRaritiesToCards maps sets to card rarities to lists of Cards.
// This map is used for generating packs.
// This map holds all cards that we can possibly generate. This should be regarded as a temporary
// solution, since we probably don't want to keep tens of thousands of cards in memory.
var SetsToRaritiesToCards = make(map[SetName]map[CardRarity][]*Card)

// GetCardsForSet checks if the set is already loaded in memory, and if it isn't, the set is read
// from the filesystem.
func GetCardsForSet(s SetName) (map[CardRarity][]*Card, error) {
	_, ok := SetsToRaritiesToCards[s]
	if !ok {
		err := ReadSetData(s)
		if err != nil {
			return nil, err
		}
	}
	m, _ := SetsToRaritiesToCards[s]
	return m, nil
}

func (c mtgjsonCard) GetCardRarity() CardRarity {
	switch {
	case c.Rarity == "Common":
		return Common
	case c.Rarity == "Uncommon":
		return Uncommon
	case c.Rarity == "Rare":
		return Rare
	case c.Rarity == "Mythic Rare":
		return Mythic
	default:
		return BasicLand
	}
}

// ReadCardData encodes json data from a byte array from a single set into a mtgjsonSet struct.
// The json data is read from the asset file on the filesystem.
// The Cards from the mtgjsonSet are stored in the SetsToRaritiesToCards global var.
func ReadSetData(setName SetName) error {
	contents, err := assets.Asset(SetNamesToJsonResources[setName])
	if err != nil {
		return err
	}

	var s mtgjsonSet
	err = json.Unmarshal(contents, &s)
	if err != nil {
		return err
	}

	raritiesToCards := make(map[CardRarity][]*Card)
	for _, card := range s.Cards {
		rarity := card.GetCardRarity()
		cardArr := raritiesToCards[rarity]
		newCard := Card{Name: card.Name, Rarity: rarity}
		raritiesToCards[rarity] = append(cardArr, &newCard)
	}

	SetsToRaritiesToCards[setName] = raritiesToCards
	return nil
}
