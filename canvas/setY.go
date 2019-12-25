package canvas

func (el *Canvas) SetY(y float64) {
	el.SelfContext.Set("y", y)
}
