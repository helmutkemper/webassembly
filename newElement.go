package iotmaker_platform_webbrowser

func NewElement() Element {
	el := Element{}
	el.InitializeDocument()

	return el
}
