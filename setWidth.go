package iotmaker_platform_webbrowser

func (el *Canvas) SetWidth(width int) {
	el.SelfContext.Set("width", width)
}
