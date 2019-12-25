package canvas

func (el *Canvas) SetPixel(x, y float64, pixel interface{}) {
	el.SelfContext.Call("putImageData", pixel, x, y)
}
