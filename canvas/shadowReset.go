package canvas

func (el *Canvas) ShadowReset() {
	el.SelfContext.Set("shadowOffsetX", 0)
	el.SelfContext.Set("shadowOffsetY", 0)
	el.SelfContext.Set("shadowBlur", 0)
	el.SelfContext.Set("shadowColor", "#000000")
}
