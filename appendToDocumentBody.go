package iotmaker_platform_webbrowser

func (el *Canvas) AppendToDocumentBody() {
	el.selfDocument.Get("body").Call("appendChild", el.SelfElement)
}
