package canvas

func (el *Canvas) SetPixel(x, y int, pixel interface{}) {
	el.SelfContext.Call("putImageData", pixel, x, y)
}
