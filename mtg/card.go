package mtg

type CardRarity int

const(
	Common		CardRarity = 0
	Uncommon	CardRarity = 1
	Rare		CardRarity = 2
	Mythic		CardRarity = 3
)

type Card struct {
	Rarity	CardRarity
	Foil	bool
}
