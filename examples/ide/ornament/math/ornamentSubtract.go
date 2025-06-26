package math

type OrnamentSubtract struct {
	OrnamentOpAmpSymbol
	Symbol  string
	AdjustX int
	AdjustY int
}

func (e *OrnamentSubtract) Init() {
	e.OrnamentOpAmpSymbol.Init()
	e.OrnamentOpAmpSymbol.SetSymbol("-")
	e.OrnamentOpAmpSymbol.SetAdjustX(0)
	e.OrnamentOpAmpSymbol.SetAdjustY(3)
}
