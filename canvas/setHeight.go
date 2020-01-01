package canvas

func (el *Canvas) SetHeight(height float64) {
	el.SelfElement.Set("height", height)
	el.SelfContext.Set("height", height)
}
