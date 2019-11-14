package iotmaker_platform_webbrowser

func NewElementById(id string) Element {
	el := Element{}
	el.InitializeById(id)

	return el
}
