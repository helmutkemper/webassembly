package components

import "github.com/helmutkemper/webassembly/browser/html"

// __passwordOnInputEvent faz a captura de dados do event input
type __passwordOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Password struct {
	__value  any
	__change *__passwordOnInputEvent

	__passwordTag *html.TagInputPassword `wasmPanel:"type:inputTagInputPassword"`
}

func (e *Password) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Password) value(value any) {
	e.__passwordTag.Value(value)
}

func (e *Password) Value(value any) {
	if e.__passwordTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
