package canvas

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext2DById(id string) {
	el.SelfElement = el.SelfElement.Call("createElement", "canvas")
	el.SelfElement.Set("id", id)
	el.SelfContextType = 1
	el.SelfContext = el.SelfElement.Call("getContext", "2d")
}
