package math

type OrnamentMultiplier struct {
	OrnamentOpAmpSymbol
}

func (o *OrnamentMultiplier) Init() {
	o.OrnamentOpAmpSymbol.Init()
	o.OrnamentOpAmpSymbol.SetSymbol("×")
	o.OrnamentOpAmpSymbol.SetAdjustX(0)
	o.OrnamentOpAmpSymbol.SetAdjustY(3)
}
