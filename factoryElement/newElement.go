package factoryElement

func NewElement() Element {
	el := Element{}
	el.InitializeDocument()

	return el
}
