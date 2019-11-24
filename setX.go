package iotmaker_platform_webbrowser

func (el *Canvas) SetX(x int) {
	el.SelfContext.Set("x", x)
}
