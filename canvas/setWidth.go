package canvas

func (el *Canvas) SetWidth(width float64) {
	el.SelfElement.Set("width", width)
	el.SelfContext.Set("width", width)
}
