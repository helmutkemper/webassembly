package canvas

func (el *Canvas) SetWidth(width int) {
	el.SelfElement.Set("width", width)
	el.SelfContext.Set("width", width)
}
