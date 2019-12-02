package canvas

func (el *Canvas) AppendToDocumentBody() {
	el.Element.Document.SelfDocument.Get("body").Call("appendChild", el.Element.SelfElement)
}
