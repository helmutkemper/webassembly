package iotmaker_platform_webbrowser

func NewExistentElementById(id string) Element {
	el := Element{}
	el.InitializeExistentElementById(id)

	return el
}
