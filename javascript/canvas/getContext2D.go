package canvas

func (el *Canvas) GetContext2D() interface{} {
	return el.SelfElement.Call("getContext", "2d")
}
