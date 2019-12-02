package canvas

func (el *Canvas) SetHeight(height int) {
	el.SelfContext.Set("height", height)
}
