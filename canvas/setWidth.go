package canvas

func (el *Canvas) SetWidth(width int) {
	el.SelfContext.Set("width", width)
}
