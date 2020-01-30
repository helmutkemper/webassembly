package canvas

func (el *Canvas) SetHeight(height int) {
	el.SelfElement.Set("height", height)
	el.SelfContext.Set("height", height)
}
