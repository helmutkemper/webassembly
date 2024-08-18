package components

import "github.com/helmutkemper/webassembly/browser/html"

type Board struct {
	__divTag *html.TagDiv
}

func (e *Board) init() {

}

func (e *Board) GetId() (id string) {
	return e.__divTag.GetId()
}
