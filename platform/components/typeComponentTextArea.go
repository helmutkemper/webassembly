package components

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// __textAreaOnInputEvent faz a captura de dados do event input
type __textAreaOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type TextArea struct {
	__value  any
	__change *__textAreaOnInputEvent

	__textAreaTag *html.TagTextArea `wasmPanel:"type:inputTagTextArea"`
}

func (e *TextArea) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *TextArea) value(value any) {
	e.__textAreaTag.Text(value)
}

func (e *TextArea) Value(value any) {
	if e.__textAreaTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
