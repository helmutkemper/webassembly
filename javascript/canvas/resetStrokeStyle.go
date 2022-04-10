package canvas

func (el *Canvas) ResetStrokeStyle() {
	el.SelfContext.Set("strokeStyle", "#000000")
}
