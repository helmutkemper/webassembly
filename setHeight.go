package iotmaker_platform_webbrowser

func (el *Canvas) SetHeight(height int) {
	el.SelfContext.Set("height", height)
}
