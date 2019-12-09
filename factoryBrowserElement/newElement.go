package factoryBrowserElement

func NewElement() Element {
	el := Element{}
	el.InitializeDocument()

	return el
}
