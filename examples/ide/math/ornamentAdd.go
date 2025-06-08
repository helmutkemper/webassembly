package math

type OrnamentAdd struct {
	OrnamentOpAmpSymbol
}

func (e *OrnamentAdd) Init() {
	e.OrnamentOpAmpSymbol.Init()
	e.OrnamentOpAmpSymbol.SetSymbol("+")
	e.OrnamentOpAmpSymbol.SetAdjustX(0)
	e.OrnamentOpAmpSymbol.SetAdjustY(4)
}
