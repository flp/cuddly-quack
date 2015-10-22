package cards

type Pack struct {
	Rare Card
	Uncommons [3]Card
	Commons [10]Card
	Foil Card
}

type PackGenerator interface {
	GetPack() Pack
}
