package element

func NewExistentElementById(id string) Element {
	el := Element{}
	el.InitializeExistentElementById(id)

	return el
}
