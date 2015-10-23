package mtg

type Pack struct {
	Cards []*Card
}

type PackGenerator interface {
	GeneratePack() Pack
}
