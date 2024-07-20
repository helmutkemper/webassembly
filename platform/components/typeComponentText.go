package components

import "github.com/helmutkemper/webassembly/browser/html"

// __textOnInputEvent faz a captura de dados do event input
type __textOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Text struct {
	__value  any
	__change *__textOnInputEvent

	__textTag *html.TagInputText `wasmPanel:"type:inputTagInputText"`
}

func (e *Text) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Text) value(value any) {
	e.__textTag.Value(value)
}

func (e *Text) Value(value any) {
	if e.__textTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
