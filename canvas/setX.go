package canvas

func (el *Canvas) SetX(x int) {
	el.SelfContext.Set("x", x)
}
