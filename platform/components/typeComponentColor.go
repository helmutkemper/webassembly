package components

import "github.com/helmutkemper/webassembly/browser/html"

// __colorOnInputEvent faz a captura de dados do event input
type __colorOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Color struct {
	__value  any
	__change *__colorOnInputEvent

	__colorTag *html.TagInputColor `wasmPanel:"type:inputTagInputColor"`
}

func (e *Color) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Color) value(value any) {
	e.__colorTag.Value(value)
}

func (e *Color) Value(value any) {
	if e.__colorTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
