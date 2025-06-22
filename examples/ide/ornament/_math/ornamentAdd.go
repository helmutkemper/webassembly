package _math

type OrnamentAdd struct {
	OrnamentOpAmpSymbol
}

func (e *OrnamentAdd) Init() (err error) {
	_ = e.OrnamentOpAmpSymbol.Init()
	e.OrnamentOpAmpSymbol.SetSymbol("+")
	e.OrnamentOpAmpSymbol.SetAdjustX(0)
	e.OrnamentOpAmpSymbol.SetAdjustY(4)
	return
}

func (e *OrnamentAdd) Update(x, y, width, height int) (err error) {
	_ = e.OrnamentOpAmpSymbol.Update(x, y, width, height)
	return
}
