package canvas

func (el *Canvas) ResetFillStyle() {
	el.SelfContext.Set("fillStyle", "#0000")
}
