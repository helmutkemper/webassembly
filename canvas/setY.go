package canvas

func (el *Canvas) SetY(y int) {
	el.SelfContext.Set("y", y)
}
