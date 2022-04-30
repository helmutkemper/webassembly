package factoryBrowserElement

func NewElement() tag.Element {
	el := tag.Element{}
	el.InitializeDocument()

	return el
}
