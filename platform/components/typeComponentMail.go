package components

import "github.com/helmutkemper/webassembly/browser/html"

// __mailOnInputEvent faz a captura de dados do event input
type __mailOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type Mail struct {
	__value  any
	__change *__mailOnInputEvent

	__mailTag *html.TagInputMail `wasmPanel:"type:inputTagInputMail"`
}

func (e *Mail) init() {
	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Mail) value(value any) {
	e.__mailTag.Value(value)
}

func (e *Mail) Value(value any) {
	if e.__mailTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
