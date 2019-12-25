package canvas

func (el *Canvas) SetX(x float64) {
	el.SelfContext.Set("x", x)
}
