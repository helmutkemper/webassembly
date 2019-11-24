package iotmaker_platform_webbrowser

func (el *Canvas) SetY(y int) {
	el.SelfContext.Set("y", y)
}
