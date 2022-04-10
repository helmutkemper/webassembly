package canvas

func (el *Canvas) ResetLineWidth() {
	el.SelfContext.Set("lineWidth", 1)
}
