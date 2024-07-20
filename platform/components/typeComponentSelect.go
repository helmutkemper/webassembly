package components

import "github.com/helmutkemper/webassembly/browser/html"

// __selectOnInputEvent faz a captura de dados do event input
type __selectOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Select struct {
	__value  any
	__change *__selectOnInputEvent

	__selectTag *html.TagSelect `wasmPanel:"type:inputTagSelect"`
}

func (e *Select) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Select) value(value any) {
	e.__selectTag.Value(value)
}

func (e *Select) Value(value any) {
	if e.__selectTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
