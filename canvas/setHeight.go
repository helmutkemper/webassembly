package canvas

func (el *Canvas) SetHeight(height float64) {
	el.SelfContext.Set("height", height)
}
