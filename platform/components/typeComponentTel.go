package components

import "github.com/helmutkemper/webassembly/browser/html"

// __telOnInputEvent faz a captura de dados do event input
type __telOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Tel struct {
	__value  any
	__change *__telOnInputEvent

	__telTag *html.TagInputTel `wasmPanel:"type:inputTagInputTel"`
}

func (e *Tel) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Tel) value(value any) {
	e.__telTag.Value(value)
}

func (e *Tel) Value(value any) {
	if e.__telTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
