package math

type OrnamentDivider struct {
	OrnamentOpAmpSymbol
	Symbol  string
	AdjustX int
	AdjustY int
}

func (e *OrnamentDivider) Init() {
	e.OrnamentOpAmpSymbol.Init()
	e.OrnamentOpAmpSymbol.SetSymbol("÷")
	e.OrnamentOpAmpSymbol.SetAdjustX(0)
	e.OrnamentOpAmpSymbol.SetAdjustY(3)
}
