package canvas

func (el *Canvas) SetWidth(width float64) {
	el.SelfContext.Set("width", width)
}
